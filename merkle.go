package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type MerkleTree struct {
	Root       *Node
	LeafHashes [][]byte
}

type Node struct {
	Left  *Node
	Right *Node
	Hash  []byte
}

func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func concatenateAndHash(left, right []byte) []byte {
	concatenatedHash := append(left, right...)
	return hashData(concatenatedHash)
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []Node

	for _, datum := range data {
		hash := hashData(datum)
		nodes = append(nodes, Node{Hash: hash})
	}

	for len(nodes) > 1 {
		var newLevel []Node

		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) {
				nodes = append(nodes, nodes[i])
			}

			left, right := nodes[i], nodes[i+1]
			newHash := concatenateAndHash(left.Hash, right.Hash)
			newLevel = append(newLevel, Node{Left: &left, Right: &right, Hash: newHash})
		}

		nodes = newLevel
	}

	tree := MerkleTree{Root: &nodes[0]}
	return &tree
}

func printTree(node *Node, level int) {
	if node == nil {
		return
	}

	fmt.Printf("%*s%s\n", level*2, "LEVEL: "+strconv.Itoa(level), " HASH: "+hex.EncodeToString(node.Hash))
	printTree(node.Left, level+1)
	printTree(node.Right, level+1)
}
