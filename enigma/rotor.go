package enigma

import (
	"fmt"
)

type Rotor struct {
	Name        string
	connections *[2][26]int
	position    int
	ringSetting int
	notches     map[int]bool
}

type RotorConfiguration struct {
	name          string
	configuration string
	position      int
	ringSetting   int
	notches       []int
}

const runeOffset = 65 // A

// NewRotor takes a configuration string of 26 characters and instantiates a rotor object
func NewRotor(r *RotorConfiguration) (*Rotor, error) {
	connections, err := convertStringConfiguration(r.configuration, r.ringSetting)
	if err != nil {
		msg := fmt.Sprintf("Unable to create rotor: %v", err)
		fmt.Println(msg)
		return nil, err
	}
	if r.position < 0 || r.position >= 26 {
		return nil, fmt.Errorf("Invalid start position %d on rotor %v", r.position, r.name)
	}
	notch := map[int]bool{}
	for _, n := range r.notches {
		if n < 0 || n >= 26 {
			return nil, fmt.Errorf("Invalid notch position %d on rotor %v", n, r.name)
		}
		notch[n] = true
	}
	return &Rotor{
		Name:        r.name,
		connections: connections,
		position:    r.position,
		ringSetting: r.ringSetting,
		notches:     notch,
	}, nil
}

// Traverse passes a signal through the rotor configuration, either forwards or backwards
func (r *Rotor) Traverse(position int, forwards bool) int {
	offsetPosition := position + r.position
	if offsetPosition > 25 {
		offsetPosition -= 26
	}
	if forwards {
		output := r.connections[0][offsetPosition] - r.position
		if output < 0 {
			output += 26
		}
		fmt.Println(position, offsetPosition, r.connections[0][offsetPosition], output, r.Name)
		return output
	}
	output := r.connections[1][offsetPosition] - r.position
	if output < 0 {
		output += 26
	}
	fmt.Println(position, offsetPosition, r.connections[1][offsetPosition], output, r.Name)
	return output
}

func (r *Rotor) IsNotchEngaged() bool {
	_, ok := r.notches[r.position]
	return ok
}

// Cycle rotates the rotor by one position
func (r *Rotor) Cycle() {
	r.position = (r.position + 1) % 26
}

// convertStringConfiguration converts a single string of characters, representing what characters [A-Z] map to
// in position 0, and returns a slice of the wire pairs
func convertStringConfiguration(conf string, ringSetting int) (*[2][26]int, error) {
	connections := emptyConnections()

	runes := []rune(conf)
	for i, r := range runes {
		if !isAllowedCharacter(r) {
			return nil, fmt.Errorf("Forbidden Character in configuration: %b", conf[i])
		}
		j := i + ringSetting
		if j > 25 {
			j -= 26
		}
		position := int(r-runeOffset) + ringSetting
		if position > 25 {
			position -= 26
		}
		connections[0][j] = position
		if connections[1][position] != -1 {
			return nil, fmt.Errorf("Duplicate Character in map, position: %v", i)
		}
		connections[1][position] = j
	}
	return connections, nil
}

func emptyConnections() *[2][26]int {
	connections := &[2][26]int{}
	for i, r := range connections {
		for j := range r {
			connections[i][j] = -1
		}
	}
	return connections
}

// isAllowedCharacter constrains the configuration string to uppercase letters
func isAllowedCharacter(r rune) bool {
	return ('A' <= r && r <= 'Z')
}

// GetRotor takes naming information for a rotor and returns the full object
func GetRotor(name string, position, setting int) (*Rotor, error) {
	n, err := getNotches(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create rotor %s: %v", name, err)
	}
	c, err := getRotor(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create rotor %s: %v", name, err)
	}
	return NewRotor(&RotorConfiguration{
		name:          name,
		configuration: c,
		position:      position,
		ringSetting:   setting,
		notches:       n,
	})
}

// GetReflector takes a reflector name and returns the configuration
func GetReflector(name string) (*Rotor, error) {
	r, err := getReflector(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create reflector %s: %v", name, err)
	}
	return NewRotor(&RotorConfiguration{
		name:          name,
		configuration: r,
		position:      0,
		ringSetting:   0,
		notches:       nil,
	})
}
