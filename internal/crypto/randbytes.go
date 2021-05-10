package crypto

import (
	"crypto/rand"
	"fmt"
	"io"
)

func randBytes(size int) []byte {
	nonce := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(fmt.Errorf("Fatal: can't retrieve rand data: %s", err.Error()))
	}
	return nonce
}
