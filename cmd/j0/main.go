package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/65535/xto17/lib/apc20"
	"github.com/65535/xto17/lib/midi"
	"github.com/65535/xto17/lib/midiout"
	"github.com/65535/xto17/lib/storage"
	"github.com/pkg/errors"
	"github.com/rakyll/portmidi"
)

var (
	// Path to the scripts and json configs.
	scriptsPath = "/Users/kereliuk/.xto17/scripts"

	// A map of each launc
	launchNumToScript = make(map[int64]string)
)

// LaunchMetadata is a config for maybe how or when launch squares light up.
type LaunchMetadata struct {
	LaunchNumber int64
}

// registerScript registers a script to the registry.
func registerScript(scriptPath string, lm LaunchMetadata) error {
	// Copy the script to the scripts folder.
	scriptDest := path.Join(scriptsPath, path.Base(scriptPath))
	err := storage.Copy(scriptPath, scriptDest)
	if err != nil {
		return errors.Wrap(err, "unable to copy script")
	}

	// Create the json metadata and save it.
	b, err := json.Marshal(lm)
	if err != nil {
		return errors.Wrap(err, "unable to marshal json of midi message")
	}
	ext := path.Ext(scriptDest)
	jsonDest := strings.TrimRight(scriptDest, ext) + ".json"
	jsonWriter, err := os.Create(jsonDest)
	if err != nil {
		return errors.Wrap(err, "unable to create json file to write to")
	}
	defer jsonWriter.Close()
	_, err = io.Copy(jsonWriter, bytes.NewReader(b))
	if err != nil {
		return errors.Wrap(err, "unable to write json bytes to file")
	}

	// Update the mapping.
	launchNumToScript[lm.LaunchNumber] = scriptDest
	return nil
}

// runScript runs an arbitrary file that it assumes to be an executable which the caller has access to execute.
func runScript(launchNumber int64, out *portmidi.Stream) error {
	fmt.Printf("Running script: %v\n\n", launchNumToScript[launchNumber])

	cmd := exec.Command(launchNumToScript[launchNumber])
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "could not open stdout pipe")
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "could not open stdout pipe")
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "could not start command")
	}
	finished := make(chan struct{})
	// Exits from an EOF
	go storage.PrintLineByLine(stdout)
	go storage.PrintLineByLine(stderr)
	// Exits by being signaled by a channel.
	go midiout.ToggleLight(launchNumber, out, finished)
	// Wait for the command to finish.
	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "could not wait on command")
	}
	close(finished)

	return nil
}

// readRegistryAndSetLaunches reads the registry and turns on the launch squares with scripts associated with them.
func readRegistryAndSetLaunches(out *portmidi.Stream) error {
	err := storage.PathIterate(scriptsPath+"/*.sh", func(filename string) error {
		ext := path.Ext(filename)
		jsonPath := strings.TrimRight(filename, ext) + ".json"
		reader, err := storage.GetPathReader(jsonPath)
		if err != nil {
			return errors.Wrap(err, "unable to open json file")
		}

		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return errors.Wrap(err, "unable to read json file")
		}

		var lm LaunchMetadata
		if err := json.Unmarshal(b, &lm); err != nil {
			return errors.Wrap(err, "unable to unmarshal json")
		}

		// Update the mapping.
		launchNumToScript[lm.LaunchNumber] = filename
		midiInfo := apc20.IntToMidiMessageDown[lm.LaunchNumber]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		fmt.Printf("Registered script: %v\n", filename)
		return nil
	})
	return err
}

func main() {
	// Initialize the lib.
	fmt.Println("Initializing...")
	portmidi.Initialize()

	// Set up the outbound stream.
	midiOut, err := portmidi.NewOutputStream(1, 1024, 0)
	if err != nil {
		log.Fatalf("could not open output stream on midi device: %v", err)
		return
	}

	// Make it look cool.
	if err := midiout.APC20Splash(midiOut); err != nil {
		log.Fatalf("could not perform the opening splash screen: %v", err)
	}

	// Input Stream.
	stream, err := portmidi.NewInputStream(0, 1024)
	if err != nil {
		log.Fatal("could not open input stream on midi device")
	}
	defer stream.Close()

	// Reload the registry.
	if err := readRegistryAndSetLaunches(midiOut); err != nil {
		log.Fatalf("could not read registry during initialization: %v", err)
	}

	// Main event loop.
	events := stream.Listen()
	fmt.Println("Listening for events...")
	for {
		// Incoming message.
		event := <-events
		eventMidiMsg := midi.MidiMessage{
			Status: event.Status,
			Data1:  event.Data1,
			Data2:  event.Data2,
		}

		// See if the message is a MIDI message for APC20.
		// If it is, execute it.
		for i := range launchNumToScript {
			if _, ok := apc20.MidiMessageUpToInt[eventMidiMsg]; !ok {
				continue
			}
			if i == apc20.MidiMessageUpToInt[eventMidiMsg] {
				err = runScript(i, midiOut)
				if err != nil {
					log.Fatalf("unable to run script: %v", err)
					return
				}
			}
		}
	}

}
