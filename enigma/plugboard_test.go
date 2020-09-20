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
			res := p.Traverse(tt.input)
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
			p, err := NewPlugboard(tt.input)
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
			name: "more than 5",
			input: [][]int{
				{0, 25},
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
			p, err := NewPlugboard(tt.input)
			assert.Nil(t, p)
			assert.Error(t, err)
		})
	}
}
