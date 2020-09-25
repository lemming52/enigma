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
		plugs     string
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
			plugs:     "",
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "space",
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
			plugs:     "",
			input:     "A AAA",
			expected:  "B DZG",
		}, {
			name: "number",
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
			plugs:     "",
			input:     "A1AAA",
			expected:  "B1DZG",
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
			plugs:     "",
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
			plugs:     "",
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
			plugs:     "",
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
			plugs:     "",
			input:     "AAAAA",
			expected:  "UBDZG",
		}, {
			name: "plugboard",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
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
			plugs:     "AZ",
			input:     "AAAZZ",
			expected:  "UTZGO",
		}, {
			name: "everything",
			rotors: []*RotorConfiguration{
				{
					name:          "VIII",
					configuration: rotorVIII,
					position:      4,
					notches:       []int{12, 25},
					ringSetting:   0,
				}, {
					name:          "II",
					configuration: rotorII,
					position:      1,
					notches:       []int{4},
					ringSetting:   4,
				}, {
					name:          "I",
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   4,
				},
			},
			reflector: reflectorC,
			plugs:     "AZ FG",
			input:     "KRKRALLEXXFOLGENDESISTSOFORTBEKANNTZUGEBENXXICHHABEFOLGELNBEBEFEHLERHALTENXXJANSTERLEDESBISHERIGXNREICHSMARSCHALLSJGOERINGJSETZTDERFUEHRERSIEYHVRRGRZSSADMIRALYALSSEINENNACHFOLGEREINXSCHRIFTLSCHEVOLLMACHTUNTERWEGSXABSOFORTSOLLENSIESAEMTLICHEMASSNAHMENVERFUEGENYDIESICHAUSDERGEGENWAERTIGENLAGEERGEBENXGEZXREICHSLEITEIKKTULPEKKJBORMANNJXXOBXDXMMMDURNHFKSTXKOMXADMXUUUBOOIEXKPO",
			expected:  "VDIZLUASOOUMUTQWXGQQNOAJLDBRYOPDBZZBDQCEICQHQZIBXWGDJJXXYERXSVUQQXYEBZFHJFEZNIIXRPJZFGVPXSJQBFKXNPGYSNZNGFCCNBBDSMICNAZVONWNVMHWRYEMAFXYBSYGMVZPIUNTFFQCIOZZXZMHOAWUNPIKOEIOIKAQCZUPAOYBFSGMSDJKQKHSDUULVGDVFMKPSQYAIGCRRSVXOZEEPWKTCHYLBRPTXUSSDXZXEQUWVMOUUKVFZUZXFJJJXMLLQDOILEUVTCYKSASRQFQIWDTBXNDMIOVSRSNEWGKBXXWPSZLYCMEPMZYJPJPZJOESEUAKTUCWHJAZDHRLHVJVSPALMMEDIIPMMRBENBMNV",
		}, {
			name: "shark backwards compatible",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
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
				}, {
					name:          "Beta",
					configuration: rotorBeta,
					position:      0,
					notches:       nil,
					ringSetting:   0,
				},
			},
			reflector: reflectorBThin,
			plugs:     "",
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "shark",
			rotors: []*RotorConfiguration{
				{
					name:          "III",
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
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
				}, {
					name:          "Gamma",
					configuration: rotorGamma,
					position:      1,
					notches:       nil,
					ringSetting:   0,
				},
			},
			reflector: reflectorCThin,
			plugs:     "",
			input:     "AAAAA",
			expected:  "NYXVI",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			e, err := New(tt.rotors, tt.reflector, tt.plugs)
			assert.Nil(t, err)
			res, err := e.Encode(tt.input)
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
			e, err := New(tt.rotors, tt.reflector, "")
			assert.Nil(t, err)
			res, err := e.Encode(tt.input)
			assert.Nil(t, err)
			for i, r := range tt.rotors {
				e.rotors[i].position = r.position
			}
			res, err = e.Encode(res)
			assert.Nil(t, err)
			assert.Equal(t, tt.input, res, "encoded string should match")
		})
	}
}
