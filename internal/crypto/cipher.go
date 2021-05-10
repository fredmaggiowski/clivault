package crypto

type Cipher interface {
	Decrypt(ciphertext []byte, username string) ([]byte, error)
	Encrypt(plaintext []byte, username string) ([]byte, error)
}
