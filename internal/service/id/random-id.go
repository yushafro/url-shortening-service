package id

import "crypto/rand"

func RandomId(length uint8) string {
	id := make([]byte, length)

	rand.Read(id)
	for i := range length {
		id[i] = CHARSET[int(id[i])%len(CHARSET)]
	}

	return string(id)
}
