package enigma

import (
	"fmt"
)

// Rotors Names of the rotors to be used to specify conveniently
const RotorI = "RotorI"
const RotorII = "RotorII"
const RotorIII = "RotorIII"
const RotorIV = "RotorIV"
const RotorV = "RotorV"
const RotorVI = "RotorVI"
const RotorVII = "RotorVII"
const RotorVIII = "RotorVIII"
const RotorBeta = "RotorBeta"
const RotorGamma = "RotorGamma"

const rotorI = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
const rotorII = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
const rotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
const rotorIV = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
const rotorV = "VZBRGITYUPSDNHLXAWMJQOFECK"
const rotorVI = "JPGVOUMFYQBENHZRDKASXLICTW"
const rotorVII = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
const rotorVIII = "FKQHTLXOCBJSPDZRAMEWNIUYGV"
const rotorBeta = "LEYJVCNIXWPBQMDRTAKZGFUHOS"
const rotorGamma = "FSOKANUERHMBTIYCWLQPZXVGJD"

func getRotor(k string) (string, error) {
	m := map[string]string{
		RotorI:     rotorI,
		RotorII:    rotorII,
		RotorIII:   rotorIII,
		RotorIV:    rotorIV,
		RotorV:     rotorV,
		RotorVI:    rotorVI,
		RotorVII:   rotorVII,
		RotorVIII:  rotorVIII,
		RotorBeta:  rotorBeta,
		RotorGamma: rotorGamma,
	}
	r, ok := m[k]
	if !ok {
		return "", fmt.Errorf("unknown rotor: %s", k)
	}
	return r, nil
}

// Notches
func getNotches(k string) ([]int, error) {
	m := map[string][]int{
		RotorI:     {16},
		RotorII:    {4},
		RotorIII:   {21},
		RotorIV:    {9},
		RotorV:     {25},
		RotorVI:    {12, 25},
		RotorVII:   {12, 25},
		RotorVIII:  {12, 25},
		RotorBeta:  nil,
		RotorGamma: nil,
	}
	n, ok := m[k]
	if !ok {
		return nil, fmt.Errorf("unknown rotor: %s", k)
	}
	return n, nil
}

// Reflectors Names of the reflectors for convenient specification
const ReflectorA = "ReflectorA"
const ReflectorB = "ReflectorB"
const ReflectorC = "ReflectorC"
const ReflectorBThin = "ReflectorBThin"
const ReflectorCThin = "ReflectorCThin"

const reflectorA = "EJMZALYXVBWFCRQUONTSPIKHGD"
const reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
const reflectorC = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
const reflectorBThin = "ENKQAUYWJICOPBLMDXZVFTHRGS"
const reflectorCThin = "RDOBJNTKVEHMLFCWZAXGYIPSUQ"

func getReflector(k string) (string, error) {
	m := map[string]string{
		ReflectorA:     reflectorA,
		ReflectorB:     reflectorB,
		ReflectorC:     reflectorC,
		ReflectorBThin: reflectorBThin,
		ReflectorCThin: reflectorCThin,
	}
	r, ok := m[k]
	if !ok {
		return "", fmt.Errorf("unknown reflector: %s", k)
	}
	return r, nil
}
