package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// User input for plaintext and key
	fmt.Print("Enter plaintext: ")
	var plaintext string
	fmt.Scanln(&plaintext)

	fmt.Print("Enter key (32 bytes): ")
	var key [32]byte
	os.Stdin.Read(key[:])

	// Create a new ChaCha20 cipher
	nonce := make([]byte, chacha20.NonceSizeX)
	c, err := chacha20.NewUnauthenticatedCipher(key[:], nonce)
	if err != nil {
		panic(err)
	}

	// Encrypt plaintext
	ciphertext := make([]byte, len(plaintext))
	c.XORKeyStream(ciphertext, []byte(plaintext))
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Decrypt ciphertext
	c, err = chacha20.NewUnauthenticatedCipher(key[:], nonce) // Recreate the cipher with the same key and nonce
	if err != nil {
		panic(err)
	}
	plaintextBytes := []byte(plaintext) // Convert plaintext to bytes
	c.XORKeyStream(plaintextBytes, ciphertext) // Decrypt directly into plaintext buffer
	fmt.Printf("Decrypted: %s\n", plaintextBytes)
}
