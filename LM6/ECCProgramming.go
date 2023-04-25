package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to encrypt: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(input)
	message := []byte(input)

	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	sharedSecretX, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	sharedSecret := sharedSecretX.Bytes()
	block, err := aes.NewCipher(sharedSecret)

	if err != nil {
		panic(err)
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, 12)

	if err != nil {
		panic(err)
	}

	ciphertext := aesGcm.Seal(nil, nonce, message, nil)

	hexCiphertext := hex.EncodeToString(ciphertext)
	fmt.Printf("Encrypted message: %s\n", hexCiphertext)

	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		panic(err)
	}
    
	fmt.Printf("Decrypted message: %s\n", plaintext)
}
