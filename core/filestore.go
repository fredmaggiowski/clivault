package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fredmaggiowski/clivault/internal/datastore"
)

type JSONFileStore struct {
	filepath string
}

func NewJSONFileStore(filepath string) *JSONFileStore {
	return &JSONFileStore{
		filepath: filepath,
	}
}

func (f *JSONFileStore) LoadBlob() (datastore.Blob, error) {
	file, err := os.OpenFile(f.filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		return nil, err
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	blob := datastore.Blob{}
	if len(fileContent) == 0 {
		return blob, nil
	}

	fmt.Printf("Before unmarshal %+v\n", fileContent)
	if err := json.Unmarshal(fileContent, &blob); err != nil {
		return nil, err
	}

	return blob, nil
}

func (f *JSONFileStore) WriteBlob(blob datastore.Blob) error {
	fileContent, err := json.Marshal(blob)
	if err != nil {
		return err
	}
	fmt.Printf("Before final write\n")

	return ioutil.WriteFile(f.filepath, fileContent, os.ModePerm)
}
