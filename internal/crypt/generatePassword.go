package crypt

import (
	"crypto/rand"
	"math/big"
)

func GenPassword(length int, types string) (string, error) {
	// Generate password with (-l, -t)
	// GenPassword need two params length and types on a plain string

	p := make([]byte, length)

	for i := 0; i < length; i++ {

		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(types))))

		if err != nil {
			return "", err
		}

		p[i] = types[num.Int64()]
	}

	return string(p), nil
}
