package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

const (
	Memory      = 64 * 1024 // 64MB
	Iterations  = 3
	Parallelism = 2
	KeyLength   = 32
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	return salt, nil
}

func GenerateKey(password string, salt []byte) ([]byte, error) {
	return argon2.IDKey(
		[]byte(password),
		salt,
		Iterations,
		Memory,
		Parallelism,
		KeyLength,
	), nil
}

func EncryptFile(plainData []byte, password string) ([]byte, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}

	key, _ := GenerateKey(password, salt)

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plainData, nil)

	return append(salt, cipherText...), nil
}

func DecryptFile(data []byte, password string) ([]byte, error) {
	salt := data[:16]
	cipherWitoutNonce := data[16:]

	key, err := GenerateKey(password, salt)
	if err != nil {
		return nil, err
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherWitoutNonce) < nonceSize {
		return nil, fmt.Errorf("Corrupt data")
	}

	nonce, ciperText := cipherWitoutNonce[:nonceSize], cipherWitoutNonce[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciperText, nil)
	if err != nil {
		return nil, fmt.Errorf("No access")
	}

	return plaintext, nil
}
