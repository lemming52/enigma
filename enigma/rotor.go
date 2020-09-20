package enigma

import (
	"fmt"
)

type Rotor struct {
	Name        string
	connections *[2][26]int
	position    int
}

type RotorConfiguration struct {
	name          string
	configuration string
	position      int
}

const runeOffset = 65 // A

// NewRotor takes a configuration string of 26 characters and instantiates a rotor object
func NewRotor(r *RotorConfiguration) (*Rotor, error) {
	connections, err := convertStringConfiguration(r.configuration)
	if err != nil {
		msg := fmt.Sprintf("Unable to create rotor: %v", err)
		fmt.Println(msg)
		return nil, err
	}
	return &Rotor{
		Name:        r.name,
		connections: connections,
		position:    r.position,
	}, nil
}

// Traverse passes a signal through the rotor configuration, either forwards or backwards
func (r *Rotor) Traverse(position int, forwards bool) int {
	if forwards {
		return r.connections[0][(position+r.position)%25]
	}
	return r.connections[1][(position+r.position)%25]
}

// Cycle rotates the rotor by one position
func (r *Rotor) Cycle() {
	r.position = (r.position + 1) % 26
}

// convertStringConfiguration converts a single string of characters, representing what characters [A-Z] map to
// in position 0, and returns a slice of the wire pairs
func convertStringConfiguration(conf string) (*[2][26]int, error) {
	connections := &[2][26]int{}
	runes := []rune(conf)
	for i, r := range runes {
		if !isAllowedCharacter(r) {
			return nil, fmt.Errorf("Forbidden Character in configuration: %b", conf[i])
		}
		position := int(r - runeOffset)
		connections[0][i] = position
		if i != 0 && connections[1][position] != 0 {
			return nil, fmt.Errorf("Duplicate Character map, characters in positions %d and %d", i, connections[1][position])
		}
		connections[1][position] = i
	}
	return connections, nil
}

// isAllowedCharacter constrains the configuration string to uppercase letters
func isAllowedCharacter(r rune) bool {
	return ('A' <= r && r <= 'Z')
}
