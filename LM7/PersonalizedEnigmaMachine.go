package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// Get user input
	fmt.Print("\nEnter plaintext: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	plaintext := []byte(strings.TrimSpace(userInput))

	// Keys and initialization vectors
	key := []byte("0123456789ABCDEF0123456789ABCDEF")
	iv := make([]byte, aes.BlockSize)
	_, _ = io.ReadFull(rand.Reader, iv)

	keyChaCha20 := make([]byte, chacha20.KeySize)
	nonce := make([]byte, chacha20.NonceSize)

	_, _ = io.ReadFull(rand.Reader, keyChaCha20)
	_, _ = io.ReadFull(rand.Reader, nonce)

	// Padding plaintext for AES-ECB
	paddedPlaintext := pkcs7Pad(plaintext, aes.BlockSize)
	fmt.Printf("\nPadded message: %x\n\n", paddedPlaintext)

	// Encrypt with AES-ECB
	ciphertext, _ := encryptECB(paddedPlaintext, key)
	fmt.Printf("Encyption: \nEncrypted with AES-ECB: %x\n", ciphertext)

	// Encrypt with AES-OFB
	ciphertext, _ = encryptOFB(ciphertext, key, iv)
	fmt.Printf("Encrypted with AES-OFB: %x\n", ciphertext)

	// Encrypt with ChaCha20
	ciphertext, _ = encryptChaCha20(ciphertext, keyChaCha20, nonce)
	fmt.Printf("Encrypted with ChaCha20: %x\n\n", ciphertext)

	// Decrypt with ChaCha20
	decrypted, _ := decryptChaCha20(ciphertext, keyChaCha20, nonce)
	fmt.Printf("Decryption: \nDecrypted with ChaCha20: %x\n", decrypted)

	// Decrypt with AES-OFB
	decrypted, _ = decryptOFB(decrypted, key, iv)
	fmt.Printf("Decrypted with AES-OFB: %x\n", decrypted)

	// Decrypt with AES-ECB
	decrypted, _ = decryptECB(decrypted, key)
	fmt.Printf("Decrypted with AES-ECB: %x\n", decrypted)

	// Unpad the decrypted message
	unpaddedDecrypted := pkcs7Unpad(decrypted, aes.BlockSize)
	fmt.Printf("\nUnpadded decrypted message: %s\n", unpaddedDecrypted)
}

// Encrypt plaintext with AES-ECB
func encryptECB(plaintext, key []byte) ([]byte, error) {

	// Create a new AES cipher with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Encrypt the padded plaintext using AES-ECB
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:], plaintext[i:])
	}

	return ciphertext, nil
}

// Decrypt ciphertext with AES-ECB
func decryptECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext using AES-ECB
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:], ciphertext[i:])
	}

	return plaintext, nil
}

// Encrypt plaintext with AES-OFB
func encryptOFB(plaintext, key, iv []byte) ([]byte, error) {

	// Create a new AES cipher with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new OFB stream cipher with the given IV
	stream := cipher.NewOFB(block, iv)

	// Encrypt the plaintext using AES-OFB
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// Decrypt ciphertext with AES-OFB
func decryptOFB(ciphertext, key, iv []byte) ([]byte, error) {

	// Decrypt the ciphertext using AES-OFB
	return encryptOFB(ciphertext, key, iv)
}

// Encrypt plaintext with ChaCha20
func encryptChaCha20(plaintext, key, nonce []byte) ([]byte, error) {

	// Create a new ChaCha20 cipher with the given key and nonce
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext using ChaCha20
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// Decrypt ciphertext with ChaCha20
func decryptChaCha20(ciphertext, key, nonce []byte) ([]byte, error) {

	// Decrypt the ciphertext using ChaCha20
	return encryptChaCha20(ciphertext, key, nonce)
}

// Pad data with PKCS#7 padding
func pkcs7Pad(data []byte, blockSize int) []byte {

	// Calculate the number of padding bytes needed
	padding := blockSize - len(data)%blockSize

	// Create a slice of padding bytes
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}

	// Append the padding bytes to the data
	return append(data, padtext...)
}

// Unpad data with PKCS#7 padding
func pkcs7Unpad(data []byte, blockSize int) []byte {

	// Get the length of the data
	length := len(data)

	// Get the number of padding bytes
	unpadding := int(data[length-1])

	// Return the unpadded data
	return data[:(length - unpadding)]
}
