package apc20

import "github.com/65535/xto17/lib/midi"

var ClipLaunchDownToInt = map[midi.MidiMessage]int64{
	// Starting from the top left, going horizontal to the right and then down at the end of the line.
	midi.MidiMessage{
		Status: 144,
		Data1:  53,
		Data2:  127,
	}: 0,
	midi.MidiMessage{
		Status: 145,
		Data1:  53,
		Data2:  127,
	}: 1,
	midi.MidiMessage{
		Status: 146,
		Data1:  53,
		Data2:  127,
	}: 2,
	midi.MidiMessage{
		Status: 147,
		Data1:  53,
		Data2:  127,
	}: 3,
	midi.MidiMessage{
		Status: 148,
		Data1:  53,
		Data2:  127,
	}: 4,
	midi.MidiMessage{
		Status: 149,
		Data1:  53,
		Data2:  127,
	}: 5,
	midi.MidiMessage{
		Status: 150,
		Data1:  53,
		Data2:  127,
	}: 6,
	midi.MidiMessage{
		Status: 151,
		Data1:  53,
		Data2:  127,
	}: 7,
	midi.MidiMessage{
		Status: 144,
		Data1:  54,
		Data2:  127,
	}: 8,
	midi.MidiMessage{
		Status: 145,
		Data1:  54,
		Data2:  127,
	}: 9,
	midi.MidiMessage{
		Status: 146,
		Data1:  54,
		Data2:  127,
	}: 10,
	midi.MidiMessage{
		Status: 147,
		Data1:  54,
		Data2:  127,
	}: 11,
	midi.MidiMessage{
		Status: 148,
		Data1:  54,
		Data2:  127,
	}: 12,
	midi.MidiMessage{
		Status: 149,
		Data1:  54,
		Data2:  127,
	}: 13,
	midi.MidiMessage{
		Status: 150,
		Data1:  54,
		Data2:  127,
	}: 14,
	midi.MidiMessage{
		Status: 151,
		Data1:  54,
		Data2:  127,
	}: 15,
	midi.MidiMessage{
		Status: 144,
		Data1:  55,
		Data2:  127,
	}: 16,
	midi.MidiMessage{
		Status: 145,
		Data1:  55,
		Data2:  127,
	}: 17,
	midi.MidiMessage{
		Status: 146,
		Data1:  55,
		Data2:  127,
	}: 18,
	midi.MidiMessage{
		Status: 147,
		Data1:  55,
		Data2:  127,
	}: 19,
	midi.MidiMessage{
		Status: 148,
		Data1:  55,
		Data2:  127,
	}: 20,
	midi.MidiMessage{
		Status: 149,
		Data1:  55,
		Data2:  127,
	}: 21,
	midi.MidiMessage{
		Status: 150,
		Data1:  55,
		Data2:  127,
	}: 22,
	midi.MidiMessage{
		Status: 151,
		Data1:  55,
		Data2:  127,
	}: 23,
	midi.MidiMessage{
		Status: 144,
		Data1:  56,
		Data2:  127,
	}: 24,
	midi.MidiMessage{
		Status: 145,
		Data1:  56,
		Data2:  127,
	}: 25,
	midi.MidiMessage{
		Status: 146,
		Data1:  56,
		Data2:  127,
	}: 26,
	midi.MidiMessage{
		Status: 147,
		Data1:  56,
		Data2:  127,
	}: 27,
	midi.MidiMessage{
		Status: 148,
		Data1:  56,
		Data2:  127,
	}: 28,
	midi.MidiMessage{
		Status: 149,
		Data1:  56,
		Data2:  127,
	}: 29,
	midi.MidiMessage{
		Status: 150,
		Data1:  56,
		Data2:  127,
	}: 30,
	midi.MidiMessage{
		Status: 151,
		Data1:  56,
		Data2:  127,
	}: 31,
	midi.MidiMessage{
		Status: 144,
		Data1:  57,
		Data2:  127,
	}: 32,
	midi.MidiMessage{
		Status: 145,
		Data1:  57,
		Data2:  127,
	}: 33,
	midi.MidiMessage{
		Status: 146,
		Data1:  57,
		Data2:  127,
	}: 34,
	midi.MidiMessage{
		Status: 147,
		Data1:  57,
		Data2:  127,
	}: 35,
	midi.MidiMessage{
		Status: 148,
		Data1:  57,
		Data2:  127,
	}: 36,
	midi.MidiMessage{
		Status: 149,
		Data1:  57,
		Data2:  127,
	}: 37,
	midi.MidiMessage{
		Status: 150,
		Data1:  57,
		Data2:  127,
	}: 38,
	midi.MidiMessage{
		Status: 151,
		Data1:  57,
		Data2:  127,
	}: 39,
}

var ClipLaunchUpToInt = map[midi.MidiMessage]int64{
	// Starting from the top left, going horizontal to the right and then down at the end of the line.
	midi.MidiMessage{
		Status: 128,
		Data1:  53,
		Data2:  127,
	}: 0,
	midi.MidiMessage{
		Status: 129,
		Data1:  53,
		Data2:  127,
	}: 1,
	midi.MidiMessage{
		Status: 130,
		Data1:  53,
		Data2:  127,
	}: 2,
	midi.MidiMessage{
		Status: 131,
		Data1:  53,
		Data2:  127,
	}: 3,
	midi.MidiMessage{
		Status: 132,
		Data1:  53,
		Data2:  127,
	}: 4,
	midi.MidiMessage{
		Status: 133,
		Data1:  53,
		Data2:  127,
	}: 5,
	midi.MidiMessage{
		Status: 134,
		Data1:  53,
		Data2:  127,
	}: 6,
	midi.MidiMessage{
		Status: 135,
		Data1:  53,
		Data2:  127,
	}: 7,
	midi.MidiMessage{
		Status: 128,
		Data1:  54,
		Data2:  127,
	}: 8,
	midi.MidiMessage{
		Status: 129,
		Data1:  54,
		Data2:  127,
	}: 9,
	midi.MidiMessage{
		Status: 130,
		Data1:  54,
		Data2:  127,
	}: 10,
	midi.MidiMessage{
		Status: 131,
		Data1:  54,
		Data2:  127,
	}: 11,
	midi.MidiMessage{
		Status: 132,
		Data1:  54,
		Data2:  127,
	}: 12,
	midi.MidiMessage{
		Status: 133,
		Data1:  54,
		Data2:  127,
	}: 13,
	midi.MidiMessage{
		Status: 134,
		Data1:  54,
		Data2:  127,
	}: 14,
	midi.MidiMessage{
		Status: 135,
		Data1:  54,
		Data2:  127,
	}: 15,
	midi.MidiMessage{
		Status: 128,
		Data1:  55,
		Data2:  127,
	}: 16,
	midi.MidiMessage{
		Status: 129,
		Data1:  55,
		Data2:  127,
	}: 17,
	midi.MidiMessage{
		Status: 130,
		Data1:  55,
		Data2:  127,
	}: 18,
	midi.MidiMessage{
		Status: 131,
		Data1:  55,
		Data2:  127,
	}: 19,
	midi.MidiMessage{
		Status: 132,
		Data1:  55,
		Data2:  127,
	}: 20,
	midi.MidiMessage{
		Status: 133,
		Data1:  55,
		Data2:  127,
	}: 21,
	midi.MidiMessage{
		Status: 134,
		Data1:  55,
		Data2:  127,
	}: 22,
	midi.MidiMessage{
		Status: 135,
		Data1:  55,
		Data2:  127,
	}: 23,
	midi.MidiMessage{
		Status: 128,
		Data1:  56,
		Data2:  127,
	}: 24,
	midi.MidiMessage{
		Status: 129,
		Data1:  56,
		Data2:  127,
	}: 25,
	midi.MidiMessage{
		Status: 130,
		Data1:  56,
		Data2:  127,
	}: 26,
	midi.MidiMessage{
		Status: 131,
		Data1:  56,
		Data2:  127,
	}: 27,
	midi.MidiMessage{
		Status: 132,
		Data1:  56,
		Data2:  127,
	}: 28,
	midi.MidiMessage{
		Status: 133,
		Data1:  56,
		Data2:  127,
	}: 29,
	midi.MidiMessage{
		Status: 134,
		Data1:  56,
		Data2:  127,
	}: 30,
	midi.MidiMessage{
		Status: 135,
		Data1:  56,
		Data2:  127,
	}: 31,
	midi.MidiMessage{
		Status: 128,
		Data1:  57,
		Data2:  127,
	}: 32,
	midi.MidiMessage{
		Status: 129,
		Data1:  57,
		Data2:  127,
	}: 33,
	midi.MidiMessage{
		Status: 130,
		Data1:  57,
		Data2:  127,
	}: 34,
	midi.MidiMessage{
		Status: 131,
		Data1:  57,
		Data2:  127,
	}: 35,
	midi.MidiMessage{
		Status: 132,
		Data1:  57,
		Data2:  127,
	}: 36,
	midi.MidiMessage{
		Status: 133,
		Data1:  57,
		Data2:  127,
	}: 37,
	midi.MidiMessage{
		Status: 134,
		Data1:  57,
		Data2:  127,
	}: 38,
	midi.MidiMessage{
		Status: 135,
		Data1:  57,
		Data2:  127,
	}: 39,
}

var IntToClipLaunchDown = map[int64]midi.MidiMessage{
	// Starting from the top left, going horizontal to the right and then down at the end of the line.
	0: midi.MidiMessage{
		Status: 144,
		Data1:  53,
		Data2:  127,
	},
	1: midi.MidiMessage{
		Status: 145,
		Data1:  53,
		Data2:  127,
	},
	2: midi.MidiMessage{
		Status: 146,
		Data1:  53,
		Data2:  127,
	},
	3: midi.MidiMessage{
		Status: 147,
		Data1:  53,
		Data2:  127,
	},
	4: midi.MidiMessage{
		Status: 148,
		Data1:  53,
		Data2:  127,
	},
	5: midi.MidiMessage{
		Status: 149,
		Data1:  53,
		Data2:  127,
	},
	6: midi.MidiMessage{
		Status: 150,
		Data1:  53,
		Data2:  127,
	},
	7: midi.MidiMessage{
		Status: 151,
		Data1:  53,
		Data2:  127,
	},
	8: midi.MidiMessage{
		Status: 144,
		Data1:  54,
		Data2:  127,
	},
	9: midi.MidiMessage{
		Status: 145,
		Data1:  54,
		Data2:  127,
	},
	10: midi.MidiMessage{
		Status: 146,
		Data1:  54,
		Data2:  127,
	},
	11: midi.MidiMessage{
		Status: 147,
		Data1:  54,
		Data2:  127,
	},
	12: midi.MidiMessage{
		Status: 148,
		Data1:  54,
		Data2:  127,
	},
	13: midi.MidiMessage{
		Status: 149,
		Data1:  54,
		Data2:  127,
	},
	14: midi.MidiMessage{
		Status: 150,
		Data1:  54,
		Data2:  127,
	},
	15: midi.MidiMessage{
		Status: 151,
		Data1:  54,
		Data2:  127,
	},
	16: midi.MidiMessage{
		Status: 144,
		Data1:  55,
		Data2:  127,
	},
	17: midi.MidiMessage{
		Status: 145,
		Data1:  55,
		Data2:  127,
	},
	18: midi.MidiMessage{
		Status: 146,
		Data1:  55,
		Data2:  127,
	},
	19: midi.MidiMessage{
		Status: 147,
		Data1:  55,
		Data2:  127,
	},
	20: midi.MidiMessage{
		Status: 148,
		Data1:  55,
		Data2:  127,
	},
	21: midi.MidiMessage{
		Status: 149,
		Data1:  55,
		Data2:  127,
	},
	22: midi.MidiMessage{
		Status: 150,
		Data1:  55,
		Data2:  127,
	},
	23: midi.MidiMessage{
		Status: 151,
		Data1:  55,
		Data2:  127,
	},
	24: midi.MidiMessage{
		Status: 144,
		Data1:  56,
		Data2:  127,
	},
	25: midi.MidiMessage{
		Status: 145,
		Data1:  56,
		Data2:  127,
	},
	26: midi.MidiMessage{
		Status: 146,
		Data1:  56,
		Data2:  127,
	},
	27: midi.MidiMessage{
		Status: 147,
		Data1:  56,
		Data2:  127,
	},
	28: midi.MidiMessage{
		Status: 148,
		Data1:  56,
		Data2:  127,
	},
	29: midi.MidiMessage{
		Status: 149,
		Data1:  56,
		Data2:  127,
	},
	30: midi.MidiMessage{
		Status: 150,
		Data1:  56,
		Data2:  127,
	},
	31: midi.MidiMessage{
		Status: 151,
		Data1:  56,
		Data2:  127,
	},
	32: midi.MidiMessage{
		Status: 144,
		Data1:  57,
		Data2:  127,
	},
	33: midi.MidiMessage{
		Status: 145,
		Data1:  57,
		Data2:  127,
	},
	34: midi.MidiMessage{
		Status: 146,
		Data1:  57,
		Data2:  127,
	},
	35: midi.MidiMessage{
		Status: 147,
		Data1:  57,
		Data2:  127,
	},
	36: midi.MidiMessage{
		Status: 148,
		Data1:  57,
		Data2:  127,
	},
	37: midi.MidiMessage{
		Status: 149,
		Data1:  57,
		Data2:  127,
	},
	38: midi.MidiMessage{
		Status: 150,
		Data1:  57,
		Data2:  127,
	},
	39: midi.MidiMessage{
		Status: 151,
		Data1:  57,
		Data2:  127,
	},
}

var IntToClipLaunchUp = map[int64]midi.MidiMessage{
	// Starting from the top left, going horizontal to the right and then down at the end of the line.
	0: midi.MidiMessage{
		Status: 128,
		Data1:  53,
		Data2:  127,
	},
	1: midi.MidiMessage{
		Status: 129,
		Data1:  53,
		Data2:  127,
	},
	2: midi.MidiMessage{
		Status: 130,
		Data1:  53,
		Data2:  127,
	},
	3: midi.MidiMessage{
		Status: 131,
		Data1:  53,
		Data2:  127,
	},
	4: midi.MidiMessage{
		Status: 132,
		Data1:  53,
		Data2:  127,
	},
	5: midi.MidiMessage{
		Status: 133,
		Data1:  53,
		Data2:  127,
	},
	6: midi.MidiMessage{
		Status: 134,
		Data1:  53,
		Data2:  127,
	},
	7: midi.MidiMessage{
		Status: 135,
		Data1:  53,
		Data2:  127,
	},
	8: midi.MidiMessage{
		Status: 128,
		Data1:  54,
		Data2:  127,
	},
	9: midi.MidiMessage{
		Status: 129,
		Data1:  54,
		Data2:  127,
	},
	10: midi.MidiMessage{
		Status: 130,
		Data1:  54,
		Data2:  127,
	},
	11: midi.MidiMessage{
		Status: 131,
		Data1:  54,
		Data2:  127,
	},
	12: midi.MidiMessage{
		Status: 132,
		Data1:  54,
		Data2:  127,
	},
	13: midi.MidiMessage{
		Status: 133,
		Data1:  54,
		Data2:  127,
	},
	14: midi.MidiMessage{
		Status: 134,
		Data1:  54,
		Data2:  127,
	},
	15: midi.MidiMessage{
		Status: 135,
		Data1:  54,
		Data2:  127,
	},
	16: midi.MidiMessage{
		Status: 128,
		Data1:  55,
		Data2:  127,
	},
	17: midi.MidiMessage{
		Status: 129,
		Data1:  55,
		Data2:  127,
	},
	18: midi.MidiMessage{
		Status: 130,
		Data1:  55,
		Data2:  127,
	},
	19: midi.MidiMessage{
		Status: 131,
		Data1:  55,
		Data2:  127,
	},
	20: midi.MidiMessage{
		Status: 132,
		Data1:  55,
		Data2:  127,
	},
	21: midi.MidiMessage{
		Status: 133,
		Data1:  55,
		Data2:  127,
	},
	22: midi.MidiMessage{
		Status: 134,
		Data1:  55,
		Data2:  127,
	},
	23: midi.MidiMessage{
		Status: 135,
		Data1:  55,
		Data2:  127,
	},
	24: midi.MidiMessage{
		Status: 128,
		Data1:  56,
		Data2:  127,
	},
	25: midi.MidiMessage{
		Status: 129,
		Data1:  56,
		Data2:  127,
	},
	26: midi.MidiMessage{
		Status: 130,
		Data1:  56,
		Data2:  127,
	},
	27: midi.MidiMessage{
		Status: 131,
		Data1:  56,
		Data2:  127,
	},
	28: midi.MidiMessage{
		Status: 132,
		Data1:  56,
		Data2:  127,
	},
	29: midi.MidiMessage{
		Status: 133,
		Data1:  56,
		Data2:  127,
	},
	30: midi.MidiMessage{
		Status: 134,
		Data1:  56,
		Data2:  127,
	},
	31: midi.MidiMessage{
		Status: 135,
		Data1:  56,
		Data2:  127,
	},
	32: midi.MidiMessage{
		Status: 128,
		Data1:  57,
		Data2:  127,
	},
	33: midi.MidiMessage{
		Status: 129,
		Data1:  57,
		Data2:  127,
	},
	34: midi.MidiMessage{
		Status: 130,
		Data1:  57,
		Data2:  127,
	},
	35: midi.MidiMessage{
		Status: 131,
		Data1:  57,
		Data2:  127,
	},
	36: midi.MidiMessage{
		Status: 132,
		Data1:  57,
		Data2:  127,
	},
	37: midi.MidiMessage{
		Status: 133,
		Data1:  57,
		Data2:  127,
	},
	38: midi.MidiMessage{
		Status: 134,
		Data1:  57,
		Data2:  127,
	},
	39: midi.MidiMessage{
		Status: 135,
		Data1:  57,
		Data2:  127,
	},
}
