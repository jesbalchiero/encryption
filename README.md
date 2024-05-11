# Caesar Cipher Encryption and Decryption Functions

## Overview
This repository contains functions for encrypting and decrypting text using the Caesar cipher. The Caesar cipher is a simple encryption technique where each letter in the text is shifted by a fixed number of positions in the alphabet.

## Features
- `encrypt(key int, plainText string) string`: This function takes a key (an integer) and the text to be encrypted. It returns the encrypted text using the Caesar cipher with the provided key.

- `decrypt(key int, encryptedText string) string`: This function takes a key (an integer) and the encrypted text. It returns the original text decrypted using the Caesar cipher with the provided key.

## Example Usage
```go

package main

import (
    "fmt"
)

func main() {
    key := 3
    plainText := "Hello, World!"

    encrypted := encrypt(key, plainText)
    fmt.Println("Encrypted text:", encrypted)

    decrypted := decrypt(key, encrypted)
    fmt.Println("Decrypted text:", decrypted)
}

```