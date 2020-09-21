package enigma

import "fmt"

type Enigma struct {
	plugs      *Plugboard
	rotors     []*Rotor
	rotorCount int
	reflector  *Rotor
}

func NewEnigma(rotorConfs []*RotorConfiguration, reflector string, plugs [][]int) (*Enigma, error) {
	p, err := NewPlugboard(plugs)
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate plugboard: %v", err)
	}
	ref, err := NewRotor(&RotorConfiguration{"reflector", reflector, 0, 0, nil})
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate reflector: %v", err)
	}
	rotorCount := len(rotorConfs)
	if rotorCount < 3 {
		return nil, fmt.Errorf("Insufficient rotors specified: %d", rotorCount)
	}
	rotors := []*Rotor{}
	for _, r := range rotorConfs {
		rot, err := NewRotor(r)
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

func (e *Enigma) Cycle() {
	cycle1, cycle2 := false, false
	if e.rotors[0].IsNotchEngaged() {
		cycle1 = true
	}
	if e.rotors[1].IsNotchEngaged() {
		cycle1 = true
		cycle2 = true
	}
	e.rotors[0].Cycle()
	if cycle1 {
		e.rotors[1].Cycle()
	}
	if cycle2 {
		e.rotors[2].Cycle()
	}
}

func (e *Enigma) Encode(r rune) rune {
	in := int(r - runeOffset)
	out := e.plugs.Traverse(in)
	e.Cycle()
	fmt.Println(e.rotors[0].position, e.rotors[1].position, e.rotors[2].position)
	for _, r := range e.rotors {
		out = r.Traverse(out, true)
	}
	out = e.reflector.Traverse(out, true)
	for i := e.rotorCount; i >= 0; i-- {
		out = e.rotors[i].Traverse(out, false)
	}
	return rune(e.plugs.Traverse(out)) + runeOffset
}

func (e *Enigma) EncodeString(s string) (string, error) {
	crib := []rune(s)
	cipher := []rune{}
	for _, c := range crib {
		fmt.Println(c, "-------")
		if !isAllowedCharacter(c) {
			return string(cipher), fmt.Errorf("unencodeable character: %v", c)
		}
		cipher = append(cipher, e.Encode(c))
	}
	return string(cipher), nil
}
