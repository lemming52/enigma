package enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		ringPosition int
		expected     *[2][26]int
	}{
		{
			name:  "base",
			input: "BCDEFGHIJKLMNOPQRSTUVWXYZA",
			expected: &[2][26]int{
				0: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0},
				1: {25, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			},
		}, {
			name:  "rotor I",
			input: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			expected: &[2][26]int{
				0: {4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
				1: {20, 22, 24, 6, 0, 3, 5, 15, 21, 25, 1, 4, 2, 10, 12, 19, 7, 23, 18, 11, 17, 8, 13, 16, 14, 9},
			},
		}, {
			name:  "reflector",
			input: "YRUHQSLDPXNGOKMIEBFZCWVJAT",
			expected: &[2][26]int{
				0: {24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
				1: {24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
			},
		}, {
			name:  "ring setting",
			input: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			expected: &[2][26]int{
				0: {15, 2, 4, 6, 8, 10, 12, 3, 16, 18, 20, 24, 22, 0, 14, 25, 5, 9, 23, 7, 1, 11, 13, 21, 19, 17},
				1: {13, 20, 1, 7, 2, 16, 3, 19, 4, 17, 5, 21, 6, 22, 14, 0, 8, 25, 9, 24, 10, 23, 12, 18, 11, 15},
			},
			ringPosition: 1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res, err := convertStringConfiguration(tt.input, tt.ringPosition)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "expected arrays should be equal")
		})
	}
}

func TestCycle(t *testing.T) {
	tests := []struct {
		name     string
		startPos int
		endPos   int
	}{
		{
			name:     "base",
			startPos: 5,
			endPos:   6,
		}, {
			name:     "overflow",
			startPos: 25,
			endPos:   0,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Rotor{position: tt.startPos}
			r.Cycle()
			assert.Equal(t, tt.endPos, r.position, "rotor did not cycle correctly, got %d expected %d", r.position, tt.endPos)
		})
	}
}

func TestTraverse(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		startPos    int
		connections *[2][26]int
		expected    int
		forwards    bool
	}{
		{
			name:     "base",
			input:    5,
			startPos: 5,
			connections: &[2][26]int{
				0: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0},
				1: {25, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			},
			expected: 6,
			forwards: true,
		},
		{
			name:     "overflow",
			input:    10,
			startPos: 20,
			connections: &[2][26]int{
				0: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0},
				1: {25, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			},
			expected: 11,
			forwards: true,
		},
		{
			name:     "backwards",
			input:    5,
			startPos: 5,
			connections: &[2][26]int{
				0: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0},
				1: {25, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			},
			expected: 4,
			forwards: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor{
				connections: tt.connections,
				position:    tt.startPos,
				ringSetting: 0,
			}
			t.Parallel()
			res := r.Traverse(tt.input, tt.forwards)
			assert.Equal(t, tt.expected, res, "output value should match expected")
		})
	}
}
