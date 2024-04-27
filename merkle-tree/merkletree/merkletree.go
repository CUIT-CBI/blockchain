package merkletree

type MerkleTree interface {
	Root() []byte
	NewNode([]byte)
	Exist([]byte) bool
	DeleteNode([]byte)
	UpdateNode(old, new []byte)
	GetProof([]byte) [][]byte
}
