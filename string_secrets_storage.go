package mtproto

import (
	"encoding/hex"
	"fmt"
)

type StringSecretsStorage struct {
	data []byte
}

func NewStringSecretsStorage(s string) (*StringSecretsStorage, error) {
	bb, err := hex.DecodeString(s)
	if err != nil {
		return &StringSecretsStorage{}, err
	}

	return &StringSecretsStorage{
		data: bb,
	}, nil
}

func (s StringSecretsStorage) Read(dst []byte) (int, error) {
	copy(dst, s.data)

	if len(dst) < len(s.data) {
		return len(dst), nil
	} else {
		return len(s.data), nil
	}
}

func (s *StringSecretsStorage) Write(p []byte) (int, error) {
	hexString := hex.EncodeToString(p)
	fmt.Println(hexString)

	return len(p), nil
}
