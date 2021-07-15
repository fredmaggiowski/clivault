package core

import (
	"fmt"

	"github.com/fredmaggiowski/clivault/internal/crypto"
	"github.com/fredmaggiowski/clivault/internal/datastore"
)

type Credentials struct {
	Username string
	Password string
}

func Save(credentials Credentials, originalBlob datastore.Blob, recordId, recordVal string) (datastore.Blob, error) {
	// TODO: load salt from previously saved data
	var salt []byte

	key := crypto.NewPBKDF2KeyGen(salt).DeriveKey(credentials.Password)

	encryptedData, err := crypto.NewAESCryptoCipher(key).Encrypt([]byte(recordVal), credentials.Username)
	if err != nil {
		return nil, err
	}

	dataRecord := GenDataRecord(BlobInfo{
		Salt: salt,
		Blob: encryptedData,
	})

	originalBlob[recordId] = dataRecord

	fmt.Println("Encrypted Data", encryptedData)
	return originalBlob, nil
}
