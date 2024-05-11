package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	plainText := "Aleatory"
	fmt.Println("Plain Text", plainText)

	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted text", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted text", decrypted)
}

// hashLetterFn rearranges letters based on the provided key
func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	// Gets the last letters based on the key
	lastLetters := string(runes[len(letter)-key:])
	// Gets the remaining letters
	leftOverLetters := string(runes[:len(letter)-key])

	return fmt.Sprintf("%s%s", lastLetters, leftOverLetters)
}

// encrypt applies the Caesar cipher encryption to the plaintext
func encrypt(key int, plainText string) (result string) {
	var encrypted strings.Builder
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			encrypted.WriteByte(byte((int(char)-'A'+key)%26 + 'A'))
		} else if char >= 'a' && char <= 'z' {
			encrypted.WriteByte(byte((int(char)-'a'+key)%26 + 'a'))
		} else {
			encrypted.WriteRune(char)
		}
	}
	return encrypted.String()
}

// decrypt applies the Caesar cipher decryption to the encrypted text
func decrypt(key int, encryptedText string) (result string) {
	var decrypted strings.Builder
	for _, char := range encryptedText {
		if char >= 'A' && char <= 'Z' {
			decrypted.WriteByte(byte((int(char)-'A'-key+26)%26 + 'A'))
		} else if char >= 'a' && char <= 'z' {
			decrypted.WriteByte(byte((int(char)-'a'-key+26)%26 + 'a'))
		} else {
			decrypted.WriteRune(char)
		}
	}
	return decrypted.String()
}
