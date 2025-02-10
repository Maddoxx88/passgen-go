package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

// Function to generate a secure password
func generatePassword(length int, useUppercase bool, useLowercase bool, useNumbers bool, useSpecial bool) (string, error) {
	// Define character sets
	var characterSet string
	if useUppercase {
		characterSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useLowercase {
		characterSet += "abcdefghijklmnopqrstuvwxyz"
	}
	if useNumbers {
		characterSet += "0123456789"
	}
	if useSpecial {
		characterSet += "!@#$%^&*"
	}

	// Ensure at least one character set is selected
	if len(characterSet) == 0 {
		return "", fmt.Errorf("no character set selected")
	}

	// Generate the password
	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, err := getRandomIndex(len(characterSet))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}
		password.WriteByte(characterSet[randomIndex])
	}

	return password.String(), nil
}

// Helper function to get a random index using a CSPRNG
func getRandomIndex(max int) (int, error) {
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	return int(randomNum.Int64()), nil
}

func main() {
	// Define command-line flags
	length := flag.Int("length", 12, "Length of the password")
	useUppercase := flag.Bool("uppercase", true, "Include uppercase letters (A-Z)")
	useLowercase := flag.Bool("lowercase", true, "Include lowercase letters (a-z)")
	useNumbers := flag.Bool("numbers", true, "Include numbers (0-9)")
	useSpecial := flag.Bool("special", true, "Include special characters (!@#$%^&*)")
	flag.Parse()

	// Generate the password
	password, err := generatePassword(*length, *useUppercase, *useLowercase, *useNumbers, *useSpecial)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println(password)
}