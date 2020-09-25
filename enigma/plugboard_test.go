package enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlugTraverse(t *testing.T) {
	tests := []struct {
		name        string
		connections map[int]int
		input       int
		expected    int
	}{
		{
			name: "base",
			connections: map[int]int{
				0:  25,
				25: 0,
			},
			input:    0,
			expected: 25,
		}, {
			name: "no plug",
			connections: map[int]int{
				0:  25,
				25: 0,
			},
			input:    1,
			expected: 1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Plugboard{connections: tt.connections}
			res := p.traverse(tt.input)
			assert.Equal(t, tt.expected, res, "")
		})
	}
}

func TestNewPlugboard(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected map[int]int
	}{
		{
			name: "base",
			input: [][]int{
				{0, 25},
			},
			expected: map[int]int{
				0:  25,
				25: 0,
			},
		}, {
			name: "multiple",
			input: [][]int{
				{0, 25},
				{7, 12},
			},
			expected: map[int]int{
				0:  25,
				25: 0,
				7:  12,
				12: 7,
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := newPlugboard(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, p.connections, "connections should match")
		})
	}
}

func TestInvalidPlugboard(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
	}{
		{
			name: "self link",
			input: [][]int{
				{0, 0},
			},
		}, {
			name: "repeated",
			input: [][]int{
				{0, 25},
				{0, 24},
			},
		}, {
			name: "repeated second",
			input: [][]int{
				{0, 25},
				{1, 25},
			},
		}, {
			name: "out of range",
			input: [][]int{
				{0, 27},
			},
		}, {
			name: "more than 10",
			input: [][]int{
				{0, 25},
				{2, 24},
				{1, 23},
				{10, 22},
				{9, 21},
				{8, 20},
				{7, 12},
				{6, 13},
				{5, 14},
				{4, 15},
				{3, 16},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := newPlugboard(tt.input)
			assert.Nil(t, p)
			assert.Error(t, err)
		})
	}
}

func TestParsePlugboard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]int
	}{
		{
			name:     "base",
			input:    "AZ",
			expected: [][]int{{0, 25}},
		}, {
			name:     "multiple",
			input:    "AZ GH",
			expected: [][]int{{0, 25}, {6, 7}},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := parseStringPlugboard(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, p, "output configuration should match")
		})
	}
}

func TestInvalidParsePlugboard(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "too many",
			input: "AZ BX CY DF GH JK QW ER TU IO PL",
		}, {
			name:  "too many characters",
			input: "AZ GHH",
		}, {
			name:  "invalid characters",
			input: "az",
		}, {
			name:  "self link",
			input: "AA",
		}, {
			name:  "repeated",
			input: "AZ AD",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := parseStringPlugboard(tt.input)
			assert.Nil(t, p)
			assert.Error(t, err)
		})
	}
}
