package enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	tests := []struct {
		name      string
		rotors    []*RotorConfiguration
		reflector string
		plugs     [][]int
		input     string
		expected  string
	}{
		{
			name: "base",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          "II",
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{21},
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "stepped",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          "II",
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			expected:  "BDZGOWCXLTKSBTMCDLPBMUQOFXYHCX",
		}, {
			name: "double stepped",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      20,
					notches:       []int{21},
				}, {
					name:          "II",
					configuration: rotorII,
					position:      3,
					notches:       []int{4},
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "AAAAA",
			expected:  "EQIBM",
		}, {
			name: "double notched",
			rotors: []*RotorConfiguration{
				{
					name:          "VIII",
					configuration: rotorVIII,
					position:      11,
					notches:       []int{12, 25},
				}, {
					name:          "II",
					configuration: rotorII,
					position:      3,
					notches:       []int{4},
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "ABCDEFGHIJKLMNO",
			expected:  "DDFQJKCQQXBZZQK",
		}, {
			name: "ring setting",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   1,
				}, {
					name:          "II",
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
					ringSetting:   0,
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   0,
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "AAAAA",
			expected:  "UBDZG",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			e, err := NewEnigma(tt.rotors, tt.reflector, nil)
			assert.Nil(t, err)
			res, err := e.EncodeString(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "encoded string should match")
		})
	}
}

func TestEncodeStringReverse(t *testing.T) {
	tests := []struct {
		name      string
		rotors    []*RotorConfiguration
		reflector string
		plugs     [][]int
		input     string
		expected  string
	}{
		{
			name: "base",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          "II",
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: reflectorB,
			plugs:     nil,
			input:     "AAAAA",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e, err := NewEnigma(tt.rotors, tt.reflector, nil)
			assert.Nil(t, err)
			res, err := e.EncodeString(tt.input)
			assert.Nil(t, err)
			for i, r := range tt.rotors {
				e.rotors[i].position = r.position
			}
			res, err = e.EncodeString(res)
			assert.Nil(t, err)
			assert.Equal(t, tt.input, res, "encoded string should match")
		})
	}
}
