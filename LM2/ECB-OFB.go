package main

import "fmt"

// Codebook maps 2-bit values to 2-bit values using XOR.
// This is used to perform simple encryption and decryption.
// Each row of the array represents a pair of plaintext/ciphertext
// values, and each column represents one of the input bits to XOR.
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}

// Message is a 4-element array of 2-bit values to be encrypted.
var message = [4]int{0b00, 0b01, 0b10, 0b11}

// IV (Initialization Vector) is a 2-bit value used to initialize the
// encryption stream for OFB mode.
var iv int = 0b10

// codebookLookup looks up a 2-bit value in the codebook and returns
// the corresponding 2-bit value.
func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	// Loop over each row of the codebook
	for i = 0; i < 4; i++ {
		// Check if the first column of the row matches the input
		if codebook[i][j] == xor {
			// If so, the second column of the row is the output
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

// codebookReverseLookup is similar to codebookLookup, but looks up
// a ciphertext value in the codebook and returns the corresponding
// plaintext value.
func codebookReverseLookup(ciphertext int) (lookupValue int) {
	var i, j int = 0, 1
	// Loop over each row of the codebook
	for i = 0; i < 4; i++ {
		// Check if the second column of the row matches the input
		if codebook[i][j] == ciphertext {
			// If so, the first column of the row is the output
			j--
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

func main() {
	// Perform ECB encryption
	fmt.Println("ECB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	ciphertext := make([]int, len(message))
	for i := 0; i < len(message); i++ {
		// Look up the ciphertext value for each plaintext value
		ciphertext[i] = codebookLookup(message[i])
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}

	// Perform ECB decryption
	fmt.Println("\nECB decryption details:")
	plaintext := make([]int, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		// Look up the plaintext value for each ciphertext value
		plaintext[i] = codebookReverseLookup(ciphertext[i])
		fmt.Printf("The deciphered value of %b is %b\n", ciphertext[i], plaintext[i])
	}

	// Perform OFB encryption
	fmt.Println("\nOFB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	stream := iv
	ciphertext = make([]int, len(message))
	for i := 0; i < len(message); i++ {
		// XOR the plaintext with the current stream
		xor := message[i] ^ stream
		// Look up the ciphertext value for the XOR result
		ciphertext[i] = codebookLookup(xor)
		// Update the stream to use as input for the next round
		stream = ciphertext[i]
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}

	// Perform OFB decryption
	fmt.Println("\nOFB decryption details:")
	stream = iv
	plaintext = make([]int, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		// XOR the ciphertext with the current stream
		xor := ciphertext[i] ^ stream
		// Look up the plaintext value for the XOR result
		plaintext[i] = codebookReverseLookup(xor)
		// Update the stream to use as input for the next round
		stream = ciphertext[i]
		fmt.Printf("The deciphered value of %b is %b\n", ciphertext[i], plaintext[i])
	}
}
