package main

import (
	"crypot/rand"
	"fmt"
	"math/big"
)



const (
	
	defaultLength = 16

	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
	symbols = "!#$%&/=?*+-_@"

)



func randomChar(charset string) (byte, error) {
	
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}

	return charset[n.Int64()], nil

}



func generatePassword(length int) (string, error) {
	
	if length < 4 {
		return "", fmt.Errorf("password length must be at least 4")
	}

	password := make([]byte, length)

	// Guarantee at least one character from each category.
	charsets := []string{
		lowercase,
		uppercase,
		numbers,
		symbols,
	}

	for i, charset := range charsets {
		char, err := randomChar(charset)
		if err != nil {
			return "", err
		}

		password[i] = char
	}

	// Combined character set for remaining characters.
	allCharacters := lowercase + uppercase + numbers + symbols

	for i := len(charsets); i < length; i++ {
		char, err := randomChar(allCharacters)
		if err != nil {
			return "", err
		}

		password[i] = char
	}

	// Shuffle the password so the first four characters aren't always lowercase, uppercase, number, symbol.
	for i := len(password) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}

		j := n.Int64()
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil

}



func main() {
	
	var length int

	fmt.Printf("Enter password length (default: %d): ", defaultLength)

	_, err := fmt.Scanln(&length)

	if err != nil {
		length = defaultLength
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Generated password:")
	fmt.Println(password)
	fmt.Println()
	fmt.Printf("Length: %d\n", len(password))

}
