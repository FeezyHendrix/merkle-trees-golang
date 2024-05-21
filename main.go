package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type ApiResponse struct {
	Endpoint string
	Data     []byte
}

func getCoins() ApiResponse {
	coins := []string{"BITCOIN", "SOLANA"}
	data, _ := json.Marshal(coins)
	return ApiResponse{Endpoint: "/get-coins", Data: data}
}

func buyCoins(successful bool) ApiResponse {
	if successful {
		status := map[string]int{"status": 100}
		data, _ := json.Marshal(status)
		return ApiResponse{Endpoint: "/buy-coins", Data: data}
	} else {
		status := map[string]int{"status": 101}
		data, _ := json.Marshal(status)
		return ApiResponse{Endpoint: "/buy-coins", Data: data}
	}
}

func detectBreakingChange(oldTree, newTree *MerkleTree) bool {
	return hex.EncodeToString(oldTree.Root.Hash) != hex.EncodeToString(newTree.Root.Hash)
}

func main() {
	// Simulate initial API calls
	initialGetCoinsResponse := getCoins()
	initialBuyCoinsResponse := buyCoins(true)

	// Hash initial responses and build Merkle tree
	initialData := [][]byte{initialGetCoinsResponse.Data, initialBuyCoinsResponse.Data}
	initialTree := NewMerkleTree(initialData)
	fmt.Println("Initial Merkle Tree:")
	printTree(initialTree.Root, 0)

	// Simulate subsequent API calls
	newGetCoinsResponse := getCoins()
	newBuyCoinsResponse := buyCoins(false)

	// Hash new responses and build new Merkle tree
	newData := [][]byte{newGetCoinsResponse.Data, newBuyCoinsResponse.Data}
	newTree := NewMerkleTree(newData)
	fmt.Println("\nNew Merkle Tree:")
	printTree(newTree.Root, 0)

	// Detect breaking changes
	if detectBreakingChange(initialTree, newTree) {
		fmt.Println("\nBreaking changes detected!")
	} else {
		fmt.Println("\nNo breaking changes detected.")
	}
}
