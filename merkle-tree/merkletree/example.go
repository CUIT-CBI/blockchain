package merkletree

import (
	"merkletree/hash"
	"merkletree/kvstore"
)

type SampleMerkleTree struct {
	db kvstore.KVStore
}

func (sample SampleMerkleTree) Root() []byte {
	// TODO
	return nil
}

func (sample SampleMerkleTree) NewNode(content []byte) {
	key := hash.Sha3Sum256(content)
	sample.db.Put(key, content)
	// TODO
}

func (sample SampleMerkleTree) Exist([]byte) bool {
	// TODO
	return false
}

func (sample SampleMerkleTree) DeleteNode([]byte) {
	// TODO
}

func (sample SampleMerkleTree) UpdateNode(old, new []byte) {
	// TODO
}

func (sample SampleMerkleTree) GetProof(content []byte) [][]byte {
	// TODO
	return nil
}
