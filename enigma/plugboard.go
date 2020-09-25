package enigma

import (
	"fmt"
	"strings"
)

// Plugboard is the internal representation of the enigma plugboard
type Plugboard struct {
	connections map[int]int
}

// newPlugboard takes int pair configurations and converts them with validation to a plugboard object
func newPlugboard(pairs [][]int) (*Plugboard, error) {
	p := Plugboard{connections: make(map[int]int)}
	if len(pairs) > 10 {
		return nil, fmt.Errorf("Too many plugs, limit is 10: %v", len(pairs))
	}
	for _, pair := range pairs {
		if len(pair) != 2 {
			return nil, fmt.Errorf("Too many entries in plugboard pair: %v", pair)
		}
		if pair[0] == pair[1] {
			return nil, fmt.Errorf("Cannot pair character with self: %v", pair)
		}
		if pair[0] < 0 || pair[0] > 25 || pair[1] < 0 || pair[1] > 25 {
			return nil, fmt.Errorf("Invalid characters: %v", pair)
		}
		_, ok1 := p.connections[pair[0]]
		_, ok2 := p.connections[pair[1]]
		if ok1 || ok2 {
			return nil, fmt.Errorf("Attempted to pair character again: %v", pair)
		}
		p.connections[pair[0]] = pair[1]
		p.connections[pair[1]] = pair[0]
	}
	return &p, nil
}

// parseStringPlugboard converts a string representation of a plugboard to a set of int pairs
func parseStringPlugboard(s string) ([][]int, error) {
	pairs := strings.Split(s, " ")
	if len(pairs) > 10 {
		return nil, fmt.Errorf("too many plugboard pairs: %d", len(pairs))
	}
	res := [][]int{}
	occupied := map[rune]bool{}
	for _, p := range pairs {
		if len(p) != 2 {
			return nil, fmt.Errorf("invalid plugboard configuration %s, can only connect two letters", p)
		}
		if p[0] == p[1] {
			return nil, fmt.Errorf("invalid plugboard connection %s, cannot connect letter to self", p)
		}
		r1 := rune(p[0])
		r2 := rune(p[1])
		if !isAllowedCharacter(r1) || !isAllowedCharacter(r2) {
			return nil, fmt.Errorf("invalid characters %s, must be upper case [A-Z]", p)
		}
		_, ok0 := occupied[r1]
		_, ok1 := occupied[r2]
		if ok0 || ok1 {
			return nil, fmt.Errorf("invalid plugboard configuration %s, repeated letter", p)
		}
		occupied[r1] = true
		occupied[r2] = true

		res = append(res, []int{int(r1 - runeOffset), int(r2 - runeOffset)})
	}
	return res, nil
}

func (p *Plugboard) traverse(input int) int {
	out, ok := p.connections[input]
	fmt.Println(input, out, ok, p.connections)
	if ok {
		return out
	}
	return input
}
