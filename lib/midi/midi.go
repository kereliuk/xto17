package midi

// MidiMessage is just portmidi.Event but dropping the timestamp. Honestly I am not sure why I didn't just pass around portmidi.Event and drop the Timestamp field whenever I don't want it.
type MidiMessage struct {
	Status int64
	Data1  int64
	Data2  int64
}
