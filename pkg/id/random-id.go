package id

import (
	"crypto/rand"
	"errors"
)

var (
	emptyIDError = errors.New("length must be > 0")
)

func RandomID(length uint8) (string, error) {
	if length == 0 {
		return "", emptyIDError
	}
	id := make([]byte, length)

	rand.Read(id)
	for i := range length {
		id[i] = Charset[int(id[i])%len(Charset)]
	}

	return string(id), nil
}
