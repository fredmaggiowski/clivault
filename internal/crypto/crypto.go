package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AESCryptoCipher struct {
	key []byte
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

	nonce := randBytes(aesgcm.NonceSize())

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
