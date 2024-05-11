package main

import (
	"fmt"
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
	var encrypted string
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			// Encrypt uppercase letter
			shifted := ((int(char) - 'A' + key) % 26) + 'A'
			encrypted += string(shifted)
		} else if char >= 'a' && char <= 'z' {
			// Encrypt lowercase letter
			shifted := ((int(char) - 'a' + key) % 26) + 'a'
			encrypted += string(shifted)
		} else {
			// Leave non-alphabetic characters unchanged
			encrypted += string(char)
		}
	}
	return encrypted
}

// decrypt applies the Caesar cipher decryption to the encrypted text
func decrypt(key int, encryptedText string) (result string) {
	var decrypted string
	for _, char := range encryptedText {
		if char >= 'A' && char <= 'Z' {
			// Decrypt uppercase letter
			shifted := ((int(char) - 'A' - key + 26) % 26) + 'A'
			decrypted += string(shifted)
		} else if char >= 'a' && char <= 'z' {
			// Decrypt lowercase letter
			shifted := ((int(char) - 'a' - key + 26) % 26) + 'a'
			decrypted += string(shifted)
		} else {
			// Leave non-alphabetic characters unchanged
			decrypted += string(char)
		}
	}
	return decrypted
}
