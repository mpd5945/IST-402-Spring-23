package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// Generate a random 256-bit key
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println("Failed to generate random key:", err)
		return
	}

	// Get user input for plaintext
	var plaintext string
	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&plaintext)

	// Convert plaintext to byte slice
	plaintextBytes := []byte(plaintext)

	// Generate a random 96-bit nonce
	nonce := make([]byte, chacha20.NonceSizeX)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Failed to generate random nonce:", err)
		return
	}

	// Create a new ChaCha20 cipher with the random key and nonce
	c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println("Failed to create ChaCha20 cipher:", err)
		return
	}

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintextBytes))
	c.XORKeyStream(ciphertext, plaintextBytes)

	// Encode the key and ciphertext as hexadecimal strings
	keyHex := hex.EncodeToString(key)
	ciphertextHex := hex.EncodeToString(ciphertext)
	fmt.Println("Key (hex):", keyHex)
	fmt.Println("Ciphertext (hex):", ciphertextHex)

	// Create a new cipher instance for decryption
	decryptionCipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println("Failed to create decryption cipher:", err)
		return
	}

	// Decrypt the ciphertext
	decrypted := make([]byte, len(ciphertext))
	decryptionCipher.XORKeyStream(decrypted, ciphertext)

	fmt.Println("Decrypted:", string(decrypted))
}
