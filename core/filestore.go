package core

type FileStore struct {
}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (f *FileStore) LoadBlob([]byte) {

}
