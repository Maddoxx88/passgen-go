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