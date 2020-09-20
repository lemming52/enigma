package enigma

import (
	"fmt"
)

type Plugboard struct {
	connections map[int]int
}

func NewPlugboard(pairs [][]int) (*Plugboard, error) {
	p := Plugboard{connections: make(map[int]int)}
	if len(pairs) > 5 {
		return nil, fmt.Errorf("Too many plugs, limit is 5: %v", len(pairs))
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

func (p *Plugboard) Traverse(input int) int {
	out, ok := p.connections[input]
	if ok {
		return out
	}
	return input
}
