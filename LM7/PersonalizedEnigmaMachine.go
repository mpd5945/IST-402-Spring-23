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

func encryptECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:], plaintext[i:])
	}

	return ciphertext, nil
}

func decryptECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Decrypt(plaintext[i:], ciphertext[i:])
	}

	return plaintext, nil
}

func encryptOFB(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func decryptOFB(ciphertext, key, iv []byte) ([]byte, error) {
	return encryptOFB(ciphertext, key, iv)
}

func encryptChaCha20(plaintext, key, nonce []byte) ([]byte, error) {
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func decryptChaCha20(ciphertext, key, nonce []byte) ([]byte, error) {
	return encryptChaCha20(ciphertext, key, nonce)
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

func pkcs7Unpad(data []byte, blockSize int) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
