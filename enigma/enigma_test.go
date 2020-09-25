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
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{21},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "space",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{21},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "A AAA",
			expected:  "B DZG",
		}, {
			name: "number",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{21},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "A1AAA",
			expected:  "B1DZG",
		}, {
			name: "stepped",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			expected:  "BDZGOWCXLTKSBTMCDLPBMUQOFXYHCX",
		}, {
			name: "double stepped",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      20,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      3,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "AAAAA",
			expected:  "EQIBM",
		}, {
			name: "double notched",
			rotors: []*RotorConfiguration{
				{
					name:          RotorVIII,
					configuration: rotorVIII,
					position:      11,
					notches:       []int{12, 25},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      3,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "ABCDEFGHIJKLMNO",
			expected:  "DDFQJKCQQXBZZQK",
		}, {
			name: "ring setting",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   1,
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
					ringSetting:   0,
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   0,
				},
			},
			reflector: ReflectorB,
			plugs:     "",
			input:     "AAAAA",
			expected:  "UBDZG",
		}, {
			name: "plugboard",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
					ringSetting:   0,
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   0,
				},
			},
			reflector: ReflectorB,
			plugs:     "AZ",
			input:     "AAAZZ",
			expected:  "UTZGO",
		}, {
			name: "everything",
			rotors: []*RotorConfiguration{
				{
					name:          RotorVIII,
					configuration: rotorVIII,
					position:      4,
					notches:       []int{12, 25},
					ringSetting:   0,
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      1,
					notches:       []int{4},
					ringSetting:   4,
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   4,
				},
			},
			reflector: ReflectorC,
			plugs:     "AZ FG",
			input:     "KRKRALLEXXFOLGENDESISTSOFORTBEKANNTZUGEBENXXICHHABEFOLGELNBEBEFEHLERHALTENXXJANSTERLEDESBISHERIGXNREICHSMARSCHALLSJGOERINGJSETZTDERFUEHRERSIEYHVRRGRZSSADMIRALYALSSEINENNACHFOLGEREINXSCHRIFTLSCHEVOLLMACHTUNTERWEGSXABSOFORTSOLLENSIESAEMTLICHEMASSNAHMENVERFUEGENYDIESICHAUSDERGEGENWAERTIGENLAGEERGEBENXGEZXREICHSLEITEIKKTULPEKKJBORMANNJXXOBXDXMMMDURNHFKSTXKOMXADMXUUUBOOIEXKPO",
			expected:  "VDIZLUASOOOMUAQWXHQQNOAJBDBRYOPDBZZBDICEICQFQZSMXWHRJJRXYERXSVHQLXYEPZGFJGEZNIIXRPJZGHVPXSJEBGKCNPHYSNQNHGCCNKBDSMIXNAZVOWWNVMFWRYEYAGPYBSYHMVJPIUBTGGQCIOZZXZMFOAWUNPIKOEIRWKACCZUPAOYBTSHZSDJKJKFSDUULVFDVGMKPSQPAIHCRRPVXOZEEPWKTCFYLBRPTXUCSDXZXEQNWVMOUUMVGBUZXGJJJXMILQDOILNUNTCYKSASRVGQIWOTBXWDMIOVURSNEWHKIXXWPSZLYCMEPMZYJPJPZJOESEUAKTUCWFJAZDFRBGVJVSPALMMEDIIPMMRBENBMNV",
		}, {
			name: "shark backwards compatible",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
					ringSetting:   0,
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   0,
				}, {
					name:          RotorBeta,
					configuration: rotorBeta,
					position:      0,
					notches:       nil,
					ringSetting:   0,
				},
			},
			reflector: ReflectorBThin,
			plugs:     "",
			input:     "AAAAA",
			expected:  "BDZGO",
		}, {
			name: "shark",
			rotors: []*RotorConfiguration{
				{
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
					ringSetting:   0,
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
					ringSetting:   0,
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
					ringSetting:   0,
				}, {
					name:          RotorGamma,
					configuration: rotorGamma,
					position:      1,
					notches:       nil,
					ringSetting:   0,
				},
			},
			reflector: ReflectorCThin,
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
					name:          RotorIII,
					configuration: rotorIII,
					position:      0,
					notches:       []int{21},
				}, {
					name:          RotorII,
					configuration: rotorII,
					position:      0,
					notches:       []int{4},
				}, {
					name:          RotorI,
					configuration: rotorI,
					position:      0,
					notches:       []int{16},
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

func TestNew(t *testing.T) {
	rotors := []struct {
		name     string
		position int
		setting  int
	}{
		{
			name:     RotorIII,
			position: 5,
			setting:  6,
		}, {
			name:     RotorII,
			position: 23,
			setting:  1,
		}, {
			name:     RotorI,
			position: 0,
			setting:  0,
		},
	}
	tests := []struct {
		name     string
		input    string
		expected string
		rotors   []struct {
			name     string
			position int
			setting  int
		}
		reflector string
		plugboard string
	}{
		{
			name:      "readme",
			input:     "AAAAA",
			expected:  "JTUJZ",
			rotors:    rotors,
			reflector: ReflectorB,
			plugboard: "AZ BC XT",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			conf := []*RotorConfiguration{}
			for _, r := range tt.rotors {
				conf = append(conf, &RotorConfiguration{name: r.name, position: r.position, ringSetting: r.setting})
			}
			e, err := New(conf, tt.reflector, tt.plugboard)
			assert.Nil(t, err)
			res, err := e.Encode(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "encoded message should match")
		})
	}

}
