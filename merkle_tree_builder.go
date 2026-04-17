package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootHash []byte
	Leaves   [][]byte
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	leaves := make([][]byte, len(data))
	for i, d := range data {
		hash := sha256.Sum256(d)
		leaves[i] = hash[:]
	}
	root := buildRoot(leaves)
	return &MerkleTree{
		RootHash: root,
		Leaves:   leaves,
	}
}

func buildRoot(leaves [][]byte) []byte {
	if len(leaves) == 1 {
		return leaves[0]
	}
	newLevel := make([][]byte, 0)
	for i := 0; i < len(leaves); i += 2 {
		var pair []byte
		if i+1 < len(leaves) {
			pair = append(leaves[i], leaves[i+1]...)
		} else {
			pair = append(leaves[i], leaves[i]...)
		}
		hash := sha256.Sum256(pair)
		newLevel = append(newLevel, hash[:])
	}
	return buildRoot(newLevel)
}

func main() {
	data := [][]byte{
		[]byte("tx1"),
		[]byte("tx2"),
		[]byte("tx3"),
	}
	tree := NewMerkleTree(data)
	fmt.Println("merkle root:", hex.EncodeToString(tree.RootHash))
}
