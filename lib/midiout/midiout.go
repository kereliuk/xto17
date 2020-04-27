package midiout

import (
	"math/rand"
	"time"

	"github.com/65535/xto17/lib/apc20"
	"github.com/pkg/errors"
	"github.com/rakyll/portmidi"
)

// APC20Splash takes a midi stream as input and then flashes the launchpad lights in some way.
func APC20Splash(out *portmidi.Stream) error {
	// Down GREEN
	for i := int64(0); i < 40; i++ {
		midiInfo := apc20.IntToMidiMessageDown[i]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Down ORANGE
	for i := int64(39); i > -1; i-- {
		midiInfo := apc20.IntToMidiMessageDown[i]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 5); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Down RED
	for _, i := range rand.Perm(40) {
		midiInfo := apc20.IntToMidiMessageDown[int64(i)]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 3); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Up
	for _, i := range rand.Perm(40) {
		midiInfo := apc20.IntToMidiMessageUp[int64(i)]
		if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
			return errors.Wrap(err, "could not write short to midi device")
		}
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}

// ToggleLight toggles the launchpad light for a given launchNumber with midi device at out output stream and the finish is signaled by the channel sygnal.
func ToggleLight(launchNumber int64, out *portmidi.Stream, signal chan struct{}) error {
	for {
		select {
		// Make sure the light is on when we finish.
		case <-signal:
			midiInfo := apc20.IntToMidiMessageDown[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, midiInfo.Data2); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			return nil
		// Run flash the red light.
		default:
			midiInfo := apc20.IntToMidiMessageDown[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 3); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			time.Sleep(250 * time.Millisecond)
			midiInfo = apc20.IntToMidiMessageUp[launchNumber]
			if err := out.WriteShort(midiInfo.Status, midiInfo.Data1, 0); err != nil {
				return errors.Wrap(err, "could not write short to midi device")
			}
			time.Sleep(250 * time.Millisecond)
		}
	}
}
