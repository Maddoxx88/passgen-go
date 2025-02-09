# Passphrase Generator

A simple and secure passphrase generator written in Go. This tool generates strong passphrases using a cryptographically secure random number generator (CSPRNG) and a customizable word list. Perfect for creating memorable yet highly secure passwords.

---

## Features
- Generates strong passphrases with high entropy.
- Uses a cryptographically secure random number generator (CSPRNG).
- Supports custom word lists (e.g., EFF Diceware, BIP-39).
- Allows customization of passphrase length and separators.
- Easy to use and extend.

---

## Prerequisites
Before running the application, ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.16 or higher recommended).

---

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Maddoxx88/passgen-go.git
   cd passgen-go

2. Download a word list (e.g., EFF Diceware or BIP-39) and save it as ``wordlist.txt`` in the project directory.:
- EFF Diceware Word List: [Download Here](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt)

- BIP-39 Word List: [Download Here](https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt)


## Usage

Running the Generator

1. Run the passphrase generator:

``` 
go run main.go
```

2. Example output:

```
Generated Passphrase: correct-horse-battery-staple-forest
```

## Customization

You can customize the passphrase generation by modifying the following parameters in the ``main.go`` file:

Word List: Replace ``wordlist.txt`` with your preferred word list.

Passphrase Length: Change the ``wordCount`` parameter (default is 5 words).

Separator: Modify the ``separator`` parameter (default is ``-``).

Example:
```
passphrase, err := generatePassphrase(wordList, 6, " ") // Generates a 6-word passphrase with spaces
```

## Building the Application

To build the application into an executable binary, run:

```
go build -o passphrase-generator
```

Then, run the generated binary:

```
./passphrase-generator
```

## Example Word Lists

Here are some popular word lists you can use:

- **EFF Diceware Word List:** Contains 7776 words, providing ~12.9 bits of entropy per word.
    - [Download EFF Diceware List](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt)

- **BIP-39 Word List:** Contains 2048 words, commonly used in cryptocurrency applications.
    - [Download BIP-39 List](https://github.com/bitcoin/bips/blob/master/bip-0039/english.txt)

## How It Works

1. **Word List**: The generator uses a predefined or custom word list to select words randomly.

2. **Randomness**: Words are selected using Go's ```crypto/rand``` package, which ensures cryptographically secure randomness.

3. Passphrase Construction: Words are combined with a separator (e.g., ```-```,``` ```, or ```_```) to create the final passphrase.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please:

Open an issue on GitHub.

Fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the ```LICENSE``` file for details.

## Acknowledgments
Inspired by the [EFF Diceware](https://www.eff.org/dice) method.

Uses Go's ```crypto/rand``` package for secure random number generation.

## Support
If you find this project useful, please give it a ⭐ on GitHub! For questions or feedback, feel free to open an issue.