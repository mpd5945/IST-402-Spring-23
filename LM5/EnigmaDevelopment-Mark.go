package main
// Plugboard represents the plugboard that swaps letters before entering rotors
type Plugboard struct {
	wiring [26]int
}

// Rotor represents a rotor with its own wiring and notch position
type Rotor struct {
	wiring [26]int
	notch int
}

// RotorSet represents a set of three rotors with their initial positions
type RotorSet struct {
	rotors [3]Rotor
	positions [3]int
}

// InputRotor represents the rotor used for the input of the Enigma machine
type InputRotor struct {
	wiring [26]int
}

// Reflector represents the fixed reflector of the Enigma machine
type Reflector struct {
	wiring [26]int
}

// EnigmaMachine represents the Enigma machine with its components
type EnigmaMachine struct {
	plugboard Plugboard
	reflector Reflector
	rotorSet RotorSet
	inputRotor InputRotor
}
