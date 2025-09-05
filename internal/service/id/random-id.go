package id

import "crypto/rand"

func RandomID(length uint8) string {
	id := make([]byte, length)

	rand.Read(id)
	for i := range length {
		id[i] = Charset[int(id[i])%len(Charset)]
	}

	return string(id)
}
