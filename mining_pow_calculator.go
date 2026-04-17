package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func mine(data string, difficulty int) string {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}
	nonce := 0
	for {
		hashStr := data + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(hashStr))
		res := hex.EncodeToString(hash[:])
		if res[:difficulty] == prefix {
			return res
		}
		nonce++
	}
}

func main() {
	start := time.Now()
	hash := mine("block-data-123", 4)
	fmt.Println("mined hash:", hash)
	fmt.Println("cost:", time.Since(start))
}
