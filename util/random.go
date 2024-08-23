package util

import (
	"crypto/rand"

	"github.com/bishalbera/modular-blockchain/types"
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() types.Hash {
	return types.HashFromBytes(RandomBytes(32))
}
