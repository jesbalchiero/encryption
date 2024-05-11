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

func hashLetterFnOriginal(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOverLetters)
}

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetters := string(runes[len(letter)-key:])     // Gets the last letters based on the key
	leftOverLetters := string(runes[:len(letter)-key]) // Gets the remaining letters

	return fmt.Sprintf("%s%s", lastLetters, leftOverLetters)
}

func encryptOriginal(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))

		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(hashLetter[letterPosition])

			return r
		}

		return r
	}

	strings.Map(findOne, plainText)
	return hashedString
}

func encrypt(key int, plainText string) (result string) {
	var encrypted string
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			shifted := ((int(char) - 'A' + key) % 26) + 'A'
			encrypted += string(shifted)
		} else if char >= 'a' && char <= 'z' {
			shifted := ((int(char) - 'a' + key) % 26) + 'a'
			encrypted += string(shifted)
		} else {
			encrypted += string(char)
		}
	}
	return encrypted
}

func decryptOriginal(key int, encryptedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""

	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))

		if pos != -1 {
			letterPosition := (pos + len(originalLetter)) % len(originalLetter)
			hashedString = hashedString + string(originalLetter[letterPosition])

			return r
		}

		return r
	}

	strings.Map(findOne, encryptedText)
	return hashedString
}

func decrypt(key int, encryptedText string) (result string) {
	var decrypted string
	for _, char := range encryptedText {
		if char >= 'A' && char <= 'Z' {
			shifted := ((int(char) - 'A' - key + 26) % 26) + 'A'
			decrypted += string(shifted)
		} else if char >= 'a' && char <= 'z' {
			shifted := ((int(char) - 'a' - key + 26) % 26) + 'a'
			decrypted += string(shifted)
		} else {
			decrypted += string(char)
		}
	}
	return decrypted
}
