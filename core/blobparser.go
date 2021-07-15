package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var (
	MagicNumber = []byte{0x63, 0x6c, 0x69, 0x76}
	Version     = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}

	ErrInvalidBlob = fmt.Errorf("invalid blob")
)

var (
	magicNumberLen = len(MagicNumber)
	intSize        = 8
)

type BlobInfo struct {
	Salt []byte
	Blob []byte
}

func ParseBlob(raw []byte) *BlobInfo {
	if len(raw) < len(MagicNumber) {
		panic(fmt.Errorf("%w: blob size is too short", ErrInvalidBlob))
	}

	magicNumber := raw[0:4]
	if bytes.Compare(magicNumber, MagicNumber) != 0 {
		panic(ErrInvalidBlob)
	}

	// TODO: whenever version will be useful 😅
	// version := raw[4:12]
	saltSizeBytes := raw[12:20]
	saltSize := binary.BigEndian.Uint64(saltSizeBytes)
	salt := raw[20 : 20+saltSize]

	dataBlob := raw[20+saltSize:]

	return &BlobInfo{
		Salt: salt,
		Blob: dataBlob,
	}
}

func WriteBlob(blob BlobInfo) []byte {
	resultBytes := MagicNumber
	resultBytes = append(resultBytes, Version...)

	saltSize := make([]byte, 8)

	binary.BigEndian.PutUint64(saltSize, uint64(len(blob.Salt)))
	resultBytes = append(resultBytes, saltSize...)

	resultBytes = append(resultBytes, blob.Salt...)

	resultBytes = append(resultBytes, blob.Blob...)

	return resultBytes
}
