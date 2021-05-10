package crypto

import (
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

const HardcodedSalt = "someHardcodedPlaintextSalt"

type PBKDF2KeyGen struct {
	salt []byte
}

func NewPBKDF2KeyGen(desiredSalt []byte) *PBKDF2KeyGen {
	salt := desiredSalt
	if salt == nil {
		salt = randBytes(10)
	}
	return &PBKDF2KeyGen{salt}
}

func (k *PBKDF2KeyGen) DeriveKey(password string) []byte {
	return pbkdf2.Key([]byte(password), k.salt, 4096, 32, sha1.New)
}
