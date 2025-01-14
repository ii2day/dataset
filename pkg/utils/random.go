package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// RandomInt64 generates a random integer.
func RandomInt64(upper ...int64) int64 { //nolint:gosec
	var innerMax int64
	if len(upper) == 0 || (len(upper) > 0 && upper[0] <= 0) {
		innerMax = 9999999999
	} else {
		innerMax = upper[0]
	}

	nBig, _ := rand.Int(rand.Reader, big.NewInt(innerMax))
	n := nBig.Int64()

	return n
}

// RandBytes generates bytes according to the given length, defaults to 32.
func RandBytes(length ...int) ([]byte, error) {
	b := make([]byte, 32)
	if len(length) != 0 {
		b = make([]byte, length[0])
	}

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// RandomHashString generates a random SHA256 string with the maximum length of 64.
func RandomHashString(length ...int) string {
	b, _ := RandBytes(1024)
	if len(length) != 0 {
		sliceLength := length[0]
		if length[0] > 64 {
			sliceLength = 64
		}
		if length[0] <= 0 {
			sliceLength = 64
		}

		return fmt.Sprintf("%x", sha256.Sum256(b))[:sliceLength]
	}

	return fmt.Sprintf("%x", sha256.Sum256(b))
}
