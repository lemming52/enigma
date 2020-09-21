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
					configuration: RotorIII,
					position:      0,
					notches:       []int{NotchIII},
				}, {
					name:          "II",
					configuration: RotorII,
					position:      0,
					notches:       []int{NotchII},
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
				},
			},
			reflector: ReflectorB,
			plugs:     nil,
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "stepped",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: RotorIII,
					position:      0,
					notches:       []int{NotchIII},
				}, {
					name:          "II",
					configuration: RotorII,
					position:      0,
					notches:       []int{NotchII},
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
				},
			},
			reflector: ReflectorB,
			plugs:     nil,
			input:     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			expected:  "BDZGOWCXLTKSBTMCDLPBMUQOFXYHCX",
		}, {
			name: "double stepped",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: RotorIII,
					position:      20,
					notches:       []int{NotchIII},
				}, {
					name:          "II",
					configuration: RotorII,
					position:      3,
					notches:       []int{NotchII},
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
				},
			},
			reflector: ReflectorB,
			plugs:     nil,
			input:     "AAAAA",
			expected:  "EQIBM",
		}, {
			name: "double notched",
			rotors: []*RotorConfiguration{
				{
					name:          "VIII",
					configuration: RotorVIII,
					position:      11,
					notches:       []int{NotchVIToVIIIA, NotchVIToVIIIB},
				}, {
					name:          "II",
					configuration: RotorII,
					position:      3,
					notches:       []int{NotchII},
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
				},
			},
			reflector: ReflectorB,
			plugs:     nil,
			input:     "ABCDEFGHIJKLMNO",
			expected:  "DDFQJKCQQXBZZQK",
		}, {
			name: "ring setting",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: RotorIII,
					position:      0,
					notches:       []int{NotchIII},
					ringSetting:   1,
				}, {
					name:          "II",
					configuration: RotorII,
					position:      0,
					notches:       []int{NotchII},
					ringSetting:   0,
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
					ringSetting:   0,
				},
			},
			reflector: ReflectorB,
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
					configuration: RotorIII,
					position:      0,
					notches:       []int{NotchIII},
				}, {
					name:          "II",
					configuration: RotorII,
					position:      0,
					notches:       []int{NotchII},
				}, {
					name:          "I",
					configuration: RotorI,
					position:      0,
					notches:       []int{NotchI},
				},
			},
			reflector: ReflectorB,
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
