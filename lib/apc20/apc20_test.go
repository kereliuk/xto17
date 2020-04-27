package apc20

import (
	"testing"
)

// Check that the maps are bijective. Don't know if this will always be required.
func TestAPCMapsAreBijective(t *testing.T) {
	for k, v := range MidiMessageDownToInt {
		// Down.
		if MidiMessageDownToInt[IntToMidiMessageDown[v]] != v {
			t.Errorf("aMaps were not inverses at %v", v)
		}
		if IntToMidiMessageDown[MidiMessageDownToInt[k]] != k {
			t.Errorf("bMaps were not inverses at %v", k)
		}
	}

	for k, v := range MidiMessageUpToInt {
		// Up.
		if MidiMessageUpToInt[IntToMidiMessageUp[v]] != v {
			t.Errorf("cMaps were not inverses at %v", v)
		}
		if IntToMidiMessageUp[MidiMessageUpToInt[k]] != k {
			t.Errorf("dMaps were not inverses at %v", v)
		}
	}
}
