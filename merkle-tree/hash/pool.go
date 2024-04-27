package hash

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

type HashPool interface {
	Get() hash.Hash
}

type DefaultHash struct{}

func (pool DefaultHash) Get() hash.Hash {
	return sha3.New256()
}

func Sha3Sum256(data []byte) []byte {
	h := sha3.New256()
	return h.Sum(data)
}
