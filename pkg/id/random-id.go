package id

import (
	"crypto/rand"
	"errors"
)

var ErrEmptyID = errors.New("length must be > 0")

func RandomID(length uint8) (string, error) {
	if length == 0 {
		return "", ErrEmptyID
	}
	urlID := make([]byte, length)

	_, err := rand.Read(urlID)
	if err != nil {
		panic(err.Error())
	}

	for i := range length {
		urlID[i] = Charset[int(urlID[i])%len(Charset)]
	}

	return string(urlID), nil
}
