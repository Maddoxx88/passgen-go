package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// Function to generate a strong passphrase
func generatePassphrase(wordCount int, separator string) (string, error) {
	// Use a predefined word list (replace with a larger list like EFF's Diceware list)
	wordList := []string{
		"apple", "banana", "carrot", "dog", "elephant", "frog", "giraffe", "house", "igloo", "jungle",
		"kangaroo", "lion", "monkey", "nest", "octopus", "penguin", "quail", "rabbit", "snake", "tiger",
		"umbrella", "vulture", "whale", "xylophone", "yak", "zebra",
	}

	// Ensure the word list is large enough (at least 7776 words for Diceware-level security)
	if len(wordList) < 7776 {
		fmt.Println("Warning: The word list is smaller than recommended for strong passphrases.")
	}

	// Generate the passphrase
	var passphrase strings.Builder
	for i := 0; i < wordCount; i++ {
		randomIndex, err := getRandomIndex(len(wordList))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}
		passphrase.WriteString(wordList[randomIndex])
		if i < wordCount-1 {
			passphrase.WriteString(separator) // Add separator between words
		}
	}

	return passphrase.String(), nil
}

// Helper function to get a random index using a CSPRNG
func getRandomIndex(max int) (int, error) {
	// Generate a cryptographically secure random number
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	return int(randomNum.Int64()), nil
}

func main() {
	// Example usage
	passphrase, err := generatePassphrase(5, "-")
	if err != nil {
		fmt.Println("Error generating passphrase:", err)
		return
	}
	fmt.Println("Generated Passphrase:", passphrase)
}