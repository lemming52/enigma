package enigma

import (
	"fmt"
	"strings"
	"unicode"
)

// Enigma is the functioning enigma machine capable of encoding a message
type Enigma struct {
	plugs      *Plugboard
	rotors     []*Rotor
	rotorCount int
	reflector  *Rotor
}

// New instantiates an enigma machine from a barebones configuration
func New(rotorConfs []*RotorConfiguration, reflector string, plugs string) (*Enigma, error) {
	for _, r := range rotorConfs {
		err := fillRotorConfiguration(r)
		if err != nil {
			return nil, err
		}
	}
	pb, err := parseStringPlugboard(plugs)
	if err != nil {
		return nil, err
	}
	p, err := newPlugboard(pb)
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate plugboard: %v", err)
	}
	ref, err := newRotor(&RotorConfiguration{"reflector", reflector, 0, 0, nil})
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate reflector: %v", err)
	}
	rotorCount := len(rotorConfs)
	if rotorCount < 3 {
		return nil, fmt.Errorf("insufficient rotors specified: %d", rotorCount)
	}
	rotors := []*Rotor{}
	for _, r := range rotorConfs {
		rot, err := newRotor(r)
		if err != nil {
			return nil, fmt.Errorf("unable to instantiate rotor: %s - %s", r.name, r.configuration)
		}
		rotors = append(rotors, rot)
	}
	return &Enigma{
		plugs:      p,
		rotors:     rotors,
		rotorCount: rotorCount - 1,
		reflector:  ref,
	}, nil
}

// cycle steps the rotors of the enigma as done by the M3; i.e. 4th or greater rotor is static and double stepping of the 2nd rotor occurs
func (e *Enigma) cycle() {
	cycle1, cycle2 := false, false
	if e.rotors[0].isNotchEngaged() {
		cycle1 = true
	}
	if e.rotors[1].isNotchEngaged() {
		cycle1 = true
		cycle2 = true
	}
	e.rotors[0].cycle()
	if cycle1 {
		e.rotors[1].cycle()
	}
	if cycle2 {
		e.rotors[2].cycle()
	}
}

func (e *Enigma) encode(r rune) rune {
	in := int(r - runeOffset)
	out := e.plugs.traverse(in)
	e.cycle()
	for _, r := range e.rotors {
		out = r.traverse(out, true)
	}
	out = e.reflector.traverse(out, true)
	for i := e.rotorCount; i >= 0; i-- {
		out = e.rotors[i].traverse(out, false)
	}
	return rune(e.plugs.traverse(out)) + runeOffset
}

// Encode is the principal method of the package, making use of the enigma machine to encode a string an cycle the machine
func (e *Enigma) Encode(s string) (string, error) {
	crib := []rune(strings.ToUpper(s))
	cipher := []rune{}
	for _, c := range crib {
		if !isAllowedCharacter(c) {
			if !isRetainCharacter(c) {
				return string(cipher), fmt.Errorf("unencodeable character: %v", c)
			}
			cipher = append(cipher, c)
			continue
		}
		cipher = append(cipher, e.encode(c))
	}
	return string(cipher), nil
}

func isRetainCharacter(r rune) bool {
	return ('0' <= r && r <= '9') || unicode.IsSpace(r)
}
