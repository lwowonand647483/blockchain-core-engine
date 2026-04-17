package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	PrevHash  string
	Hash      string
	Data      string
	Timestamp int64
	Nonce     int
}

func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.PrevHash + b.Data + strconv.FormatInt(b.Timestamp, 10) + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func ValidateBlock(block Block, prev Block) bool {
	if block.PrevHash != prev.Hash {
		return false
	}
	if calculateHash(block) != block.Hash {
		return false
	}
	if block.Index != prev.Index+1 {
		return false
	}
	return true
}

func main() {
	genesis := Block{
		Index:     0,
		PrevHash:  "0",
		Data:      "genesis block",
		Timestamp: time.Now().Unix(),
		Nonce:     0,
	}
	genesis.Hash = calculateHash(genesis)

	block1 := Block{
		Index:     1,
		PrevHash:  genesis.Hash,
		Data:      "transfer 10",
		Timestamp: time.Now().Unix(),
		Nonce:     12345,
	}
	block1.Hash = calculateHash(block1)

	fmt.Println("valid:", ValidateBlock(block1, genesis))
}
