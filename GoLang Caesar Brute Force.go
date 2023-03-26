package main

import (
	"fmt"
	"strings"
)

func shiftChar(ch rune, shift int) rune {
	if ch >= 'a' && ch <= 'z' {
		return 'a' + ((ch - 'a' - rune(shift) + 26) % 26)
	} else if ch >= 'A' && ch <= 'Z' {
		return 'A' + ((ch - 'A' - rune(shift) + 26) % 26)
	} else {
		return ch
	}
}

func bruteForce(cipherText string) {
	for i := 1; i <= 99; i++ { 
		var plainText strings.Builder
		for _, ch := range cipherText {
			plainText.WriteRune(shiftChar(ch, i))
		}
		fmt.Printf("Key %d: %s\n", i, plainText.String())
	}
}

func main() {
	cipherText := "Khoor Zruog!"
	bruteForce(cipherText)
}
