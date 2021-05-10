package crypto

import (
	"encoding/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AESCryptoCipher", func() {
	Context("Encrypt and Decrypt", func() {
		It("properly encrypts and decrypts a plaintext", func() {
			key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
			plaintext := []byte("exampleplaintext")
			username := "someuser"

			cipher := NewAESCryptoCipher(key)
			Expect(cipher).To(Not(BeNil()))

			result, err := cipher.Encrypt(plaintext, username)
			Expect(err).To(BeNil())

			decryptedPlaintext, err := cipher.Decrypt(result, username)
			Expect(err).To(BeNil())

			Expect(decryptedPlaintext).To(Equal(plaintext))
		})
	})
})
