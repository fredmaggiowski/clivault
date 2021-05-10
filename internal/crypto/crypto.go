package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type AESCryptoCipher struct {
	key []byte
}

func makeNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func NewAESCryptoCipher(key []byte) *AESCryptoCipher {
	return &AESCryptoCipher{key}
}

func (aesc *AESCryptoCipher) Encrypt(plaintext []byte, username string) ([]byte, error) {
	block, err := aes.NewCipher(aesc.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, err := makeNonce(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nonce, nonce, []byte(plaintext), []byte(username)), nil
}

func (aesc *AESCryptoCipher) Decrypt(ciphertext []byte, username string) ([]byte, error) {
	block, err := aes.NewCipher(aesc.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("invalid ciphertext size")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, []byte(username))
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
