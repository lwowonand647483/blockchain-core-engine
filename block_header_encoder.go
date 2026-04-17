package main

import (
	"encoding/binary"
	"fmt"
	"bytes"
)

type BlockHeader struct {
	Version   uint32
	PrevHash  [32]byte
	MerkleRoot [32]byte
	Timestamp uint64
	Difficulty uint32
	Nonce     uint64
}

func Encode(header BlockHeader) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, header.Version)
	buf.Write(header.PrevHash[:])
	buf.Write(header.MerkleRoot[:])
	binary.Write(buf, binary.BigEndian, header.Timestamp)
	binary.Write(buf, binary.BigEndian, header.Difficulty)
	binary.Write(buf, binary.BigEndian, header.Nonce)
	return buf.Bytes()
}

func main() {
	header := BlockHeader{Version: 1}
	data := Encode(header)
	fmt.Println("encoded length:", len(data))
}
