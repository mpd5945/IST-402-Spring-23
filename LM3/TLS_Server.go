package main

import (
	"crypto/cipher"
	"crypto/aes"
	"crypto/rand"
	"crypto/tls"
	"os"
	"log"
	"fmt"
	"io"
	"net"
	"bufio"
	
)

func main() {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Fatalf("Failed to conncet: %v", err)
	}

	defer conn.Close()

	fmt.Println("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	//Encrypt
	ciphertext, err := encrypt(input, conn)
	if err != nil {
		log.Fatalf("Failed to encrypt: %v", err)
	}
	fmt.Printf("ciphered Text: %v\n", ciphertext)

	//Decrypt
	plaintext, err := decrypt(ciphertext, conn)
	if err != nil {
		log.Fatalf("Failed to encrypt: %v", err)
	}
	fmt.Printf("Plain Text: %v\n", plaintext)
}

func encrypt(plaintext string, conn net.Conn) ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate ket: %v", err)
	}

	_, err = conn.Write(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to write key: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to create cipher: %v", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("Failed to genereate IV: %v", err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return ciphertext, nil;
}
	
func decrypt(ciphertext []byte, conn net.Conn) (string, error) {
	key := make([]byte, 32)
	_, err := io.ReadFull(conn, key)
	if err != nil {
		return "", fmt.Errorf("Failed to create cipher: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("Failed to create cipher: %v", err)
	}

	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	iv:= ciphertext[:aes.BlockSize]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize: ])

	return string(plaintext), nil
}