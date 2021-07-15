package core

type Credentials struct {
	Username string
	Password string
}

func Save(credentials Credentials, blob []byte, recordId, recordVal string) ([]byte, error) {
	return nil, nil
	// 	// TODO: load salt from previously saved data
	// 	var salt []byte

	// 	key := crypto.NewPBKDF2KeyGen(salt).DeriveKey(credentials.Password)

	// 	crypto.NewAESCryptoCipher(key).Encrypt
}
