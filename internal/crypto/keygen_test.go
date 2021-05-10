package crypto

import (
	"encoding/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PBKDF2KeyGen", func() {
	Context("NewPBKDF2KeyGen", func() {
		It("uses provided salt", func() {
			salt := []byte("MySalt")
			generator := NewPBKDF2KeyGen(salt)

			Expect(generator.salt).To(Equal(salt))
		})
	})

	Context("DeriveKey", func() {
		It("properly derives a key", func() {
			salt := []byte("MyVeryStrongSalt")
			password := "password"
			expectedKey, _ := hex.DecodeString("61756c9ad62a7fd3f5cdae6d5e104568b7612267ce136b1456b09009a192c747")

			generator := NewPBKDF2KeyGen(salt)
			key := generator.DeriveKey(password)
			Expect(key).To(Equal(expectedKey))
		})
	})
})
