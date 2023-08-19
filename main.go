package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

// Static key for encryption and decryption (AES-256 requires a 32-byte key)
var key = []byte("ThisIsASecretKey1234567890123456")

func encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(cipherText string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./main crypt <password> or ./main decrypt <encrypted_password>")
		os.Exit(1)
	}

	if os.Args[1] == "crypt" {
		encrypted, err := encrypt(os.Args[2], key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			os.Exit(1)
		}
		err = os.WriteFile("pwd.txt", []byte(encrypted), 0644)
		if err != nil {
			fmt.Println("Error writing to pwd.txt:", err)
			os.Exit(1)
		}
		fmt.Println("Password encrypted and saved to pwd.txt")
	} else if os.Args[1] == "decrypt" {
		cipherText, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("Error reading pwd.txt:", err)
			os.Exit(1)
		}
		decrypted, err := decrypt(strings.TrimSpace(string(cipherText)), key)
		if err != nil {
			fmt.Println("Decryption error:", err)
			os.Exit(1)
		}
		fmt.Println("Decrypted Password:", decrypted)
	} else {
		fmt.Println("Usage: ./main crypt <password> or ./main decrypt <encrypted_password>")
		os.Exit(1)
	}
}
