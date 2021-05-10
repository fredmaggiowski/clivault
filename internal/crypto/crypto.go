package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
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

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), []byte(username))
	// fmt.Printf("DATA USED IN ENC %v %s %v\n", nonce, username, ciphertext)
	return ciphertext, nil
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
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// fmt.Printf("DATA USED IN DEC %v %s\n", nonce, username)
	plaintext, err := gcm.Open(nil, nonce, ciphertext, []byte(username))
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
