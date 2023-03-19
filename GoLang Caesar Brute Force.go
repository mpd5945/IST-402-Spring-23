package main

import (
	"fmt"
	"strings"
)

func cipher(n int, plaintext string) string {
	ALPHABET := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder
	for _, l := range plaintext {
		if strings.ContainsRune(ALPHABET, l) {
			index := strings.IndexRune(ALPHABET, l)
			i := (index + n) % 26
			result.WriteString(ALPHABET[i : i+1])
		} else if strings.ContainsRune(alphabet, l) {
			index := strings.IndexRune(alphabet, l)
			i := (index + n) % 26
			result.WriteString(alphabet[i : i+1])
		} else {
			result.WriteString(string(l))
		}
	}
	return result.String()
}

func decipher(n int, cipheredtext string) string {
	ALPHABET := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder
	for _, l := range cipheredtext {
		if strings.ContainsRune(ALPHABET, l) {
			index := strings.IndexRune(ALPHABET, l)
			i := (index - n + 26) % 26
			result.WriteString(ALPHABET[i : i+1])
		} else if strings.ContainsRune(alphabet, l) {
			index := strings.IndexRune(alphabet, l)
			i := (index - n + 26) % 26
			result.WriteString(alphabet[i : i+1])
		} else {
			result.WriteString(string(l))
		}
	}
	return result.String()
}

func main() {
	message := "Why did the Go programmer quit his job? He didn't get arrays."
	key := 17
	enc := cipher(key, message)
	fmt.Printf("%d . %s\n", key, enc)

	dec := decipher(key, enc)
	fmt.Printf("%d . %s\n", key, dec)
}
