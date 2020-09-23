package enigma

import (
	"fmt"
)

// Rotors
const RotorI = "RotorI"
const RotorII = "RotorII"
const RotorIII = "RotorIII"
const RotorIV = "RotorIV"
const RotorV = "RotorV"
const RotorVI = "RotorVI"
const RotorVII = "RotorVII"
const RotorVIII = "RotorVIII"

const rotorI = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
const rotorII = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
const rotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
const rotorIV = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
const rotorV = "VZBRGITYUPSDNHLXAWMJQOFECK"
const rotorVI = "JPGVOUMFYQBENHZRDKASXLICTW"
const rotorVII = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
const rotorVIII = "FKQHTLXOCBJSPDZRAMEWNIUYGV"

func getRotor(k string) (string, error) {
	m := map[string]string{
		RotorI:    rotorI,
		RotorII:   rotorII,
		RotorIII:  rotorIII,
		RotorIV:   rotorIV,
		RotorV:    rotorV,
		RotorVI:   rotorVI,
		RotorVII:  rotorVII,
		RotorVIII: rotorVIII,
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
		RotorI:    {16},
		RotorII:   {4},
		RotorIII:  {21},
		RotorIV:   {9},
		RotorV:    {25},
		RotorVI:   {12, 25},
		RotorVII:  {12, 25},
		RotorVIII: {12, 25},
	}
	n, ok := m[k]
	if !ok {
		return nil, fmt.Errorf("unknown rotor: %s", k)
	}
	return n, nil
}

// Reflectors
const ReflectorA = "ReflectorA"
const ReflectorB = "ReflectorB"
const ReflectorC = "ReflectorC"

const reflectorA = "EJMZALYXVBWFCRQUONTSPIKHGD"
const reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
const reflectorC = "FVPJIAOYEDRZXWGCTKUQSBNMHL"

func getReflector(k string) (string, error) {
	m := map[string]string{
		ReflectorA: reflectorA,
		ReflectorB: reflectorB,
		ReflectorC: reflectorC,
	}
	r, ok := m[k]
	if !ok {
		return "", fmt.Errorf("unknown reflector: %s", k)
	}
	return r, nil
}
