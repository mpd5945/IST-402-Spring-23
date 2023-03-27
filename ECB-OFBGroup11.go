package main

import "fmt"

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var iv int = 0b10

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

func main() {
	// ECB encryption
	fmt.Println("ECB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	ciphertext := make([]int, len(message))
	for i := 0; i < len(message); i++ {
		ciphertext[i] = codebookLookup(message[i])
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}

	// OFB encryption
	fmt.Println("\nOFB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	stream := iv
	ciphertext = make([]int, len(message))
	for i := 0; i < len(message); i++ {
		xor := message[i] ^ stream
		ciphertext[i] = codebookLookup(xor)
		stream = ciphertext[i]
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}
}
