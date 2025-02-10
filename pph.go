package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
	"unicode"
)

// Function to load the word list from a file
func loadWordList(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read word list file: %v", err)
	}

	// Split the file content into lines
	lines := strings.Split(string(data), "\n")

	// Extract words from lines
	wordList := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing whitespace
		if line != "" {                // Skip empty lines
			wordList = append(wordList, line)
		}
	}

	if len(wordList) == 0 {
		return nil, fmt.Errorf("word list is empty")
	}

	return wordList, nil
}

// Function to generate a strong passphrase
func generatePassphrase(wordList []string, wordCount int, separator string, capitalize bool, includeNumber bool) (string, error) {
	if len(wordList) == 0 {
		return "", fmt.Errorf("word list is empty")
	}

	var passphrase strings.Builder
	for i := 0; i < wordCount; i++ {
		randomIndex, err := getRandomIndex(len(wordList))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %v", err)
		}

		word := wordList[randomIndex]
		if capitalize {
			word = capitalizeWord(word) // Capitalize the word
		}
		passphrase.WriteString(word)

		if i < wordCount-1 {
			passphrase.WriteString(separator)
		}
	}

	// Add a random number if requested
	if includeNumber {
		randomNum, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}
		passphrase.WriteString(separator)
		passphrase.WriteString(randomNum.String())
	}

	return passphrase.String(), nil
}

// Helper function to capitalize a word
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return string(unicode.ToUpper(rune(word[0]))) + word[1:]
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
	wordCount := flag.Int("words", 3, "Number of words in the passphrase")
	separator := flag.String("separator", "-", "Separator between words")
	capitalize := flag.Bool("capitalize", false, "Capitalize each word")
	includeNumber := flag.Bool("number", false, "Include a random number in the passphrase")
	flag.Parse()

	// Load the word list from a file
	wordList, err := loadWordList("./wordlist.txt")
	if err != nil {
		fmt.Println("Error loading word list:", err)
		return
	}

	// Generate the passphrase
	passphrase, err := generatePassphrase(wordList, *wordCount, *separator, *capitalize, *includeNumber)
	if err != nil {
		fmt.Println("Error generating passphrase:", err)
		return
	}
	fmt.Println(passphrase)
}

