package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/65535/xto17/lib/apc20"
	"github.com/65535/xto17/lib/midi"
	"github.com/pkg/errors"
	"github.com/rakyll/portmidi"
)

var (
	scriptsPath  = "/Users/kereliuk/.xto17/scripts"
	clipToScript = make(map[int64]string)
)

type LaunchMetadata struct {
	LaunchNumber int64
}

// Copy stolen from stackoverflow
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func GetPathReader(input string) (io.Reader, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func PathIterate(glob string, cb func(string) error) error {
	fileNames, err := filepath.Glob(glob)
	if err != nil {
		return errors.Wrap(err, "failed to read dir")
	}

	for _, f := range fileNames {
		fileInfo, err := os.Stat(f)
		if err != nil {
			return errors.Wrap(err, "failed to stat path")
		}

		if fileInfo.IsDir() {
			continue
		}

		if err := cb(f); err != nil {
			return errors.Wrap(err, "failed callback")
		}
	}

	return nil
}

func ReadRegistry(out *portmidi.Stream) error {
	err := PathIterate(scriptsPath+"/*.sh", func(filename string) error {
		ext := path.Ext(filename)
		jsonPath := strings.TrimRight(filename, ext) + ".json"
		reader, err := GetPathReader(jsonPath)
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
		clipToScript[lm.LaunchNumber] = filename
		midiInfo := apc20.IntToClipLaunchDown[lm.LaunchNumber]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		fmt.Printf("Registered script: %v\n", filename)
		return nil
	})
	return err
}

// RegisterScript registers a script to the config.
func RegisterScript(scriptPath string, lm LaunchMetadata) error {
	// Copy the script to the scripts folder.
	scriptDest := path.Join(scriptsPath, path.Base(scriptPath))
	err := Copy(scriptPath, scriptDest)
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
	clipToScript[lm.LaunchNumber] = scriptDest
	return nil
}

func PrintLineByLine(stdout io.ReadCloser) {
	r := bufio.NewReader(stdout)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERR: %v", err)
			break
		}
		fmt.Printf("%s\n", string(line))
	}
}

func ToggleLight(launchNumber int64, out *portmidi.Stream, signal chan struct{}) error {
	for {
		select {
		case <-signal:
			// Make sure the light is on when we finish.
			midiInfo := apc20.IntToClipLaunchDown[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			return nil
		default:
			midiInfo := apc20.IntToClipLaunchDown[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 3); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			time.Sleep(250 * time.Millisecond)
			midiInfo = apc20.IntToClipLaunchUp[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 0); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func RunScript(launchNumber int64, out *portmidi.Stream) error {
	fmt.Printf("Running script: %v\n\n", clipToScript[launchNumber])

	cmd := exec.Command(clipToScript[launchNumber])
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
	go PrintLineByLine(stdout)
	go PrintLineByLine(stderr)
	// Exits by being signaled by a channel.
	go ToggleLight(launchNumber, out, finished)
	// Wait for the command to finish.
	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "could not wait on command")
	}
	close(finished)

	return nil
}

func APC20Splash(out *portmidi.Stream) error {
	// Down GREEN
	for i := int64(0); i < 40; i++ {
		midiInfo := apc20.IntToClipLaunchDown[i]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Down ORANGE
	for i := int64(39); i > -1; i-- {
		midiInfo := apc20.IntToClipLaunchDown[i]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 5); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Down RED
	for _, i := range rand.Perm(40) {
		midiInfo := apc20.IntToClipLaunchDown[int64(i)]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 3); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Up
	for _, i := range rand.Perm(40) {
		midiInfo := apc20.IntToClipLaunchUp[int64(i)]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	return nil
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
	if err := APC20Splash(midiOut); err != nil {
		log.Fatalf("could not perform the opening splash screen: %v", err)
	}

	// Input Stream.
	stream, err := portmidi.NewInputStream(0, 1024)
	if err != nil {
		log.Fatal("could not open input stream on midi device")
	}
	defer stream.Close()

	// Reload the registry.
	if err := ReadRegistry(midiOut); err != nil {
		log.Fatalf("could not read registry during initialization: %v", err)
	}

	events := stream.Listen()
	fmt.Println("Listening for events...")
	for {
		event := <-events
		eventMidiMsg := midi.MidiMessage{
			Status: event.Status,
			Data1:  event.Data1,
			Data2:  event.Data2,
		}

		for i := range clipToScript {
			if _, ok := apc20.ClipLaunchUpToInt[eventMidiMsg]; !ok {
				continue
			}
			if i == apc20.ClipLaunchUpToInt[eventMidiMsg] {
				err = RunScript(i, midiOut)
				if err != nil {
					log.Fatalf("unable to run script: %v", err)
					return
				}
			}
		}
	}

}
