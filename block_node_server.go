package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"sync"
	"time"
)

type ChainNode struct {
	NodeID    string
	Addr      string
	Connected []string
	Lock      sync.Mutex
}

func NewNode(addr string) *ChainNode {
	hash := sha256.Sum256([]byte(addr + time.Now().String()))
	return &ChainNode{
		NodeID:    hex.EncodeToString(hash[:]),
		Addr:      addr,
		Connected: make([]string, 0),
	}
}

func (n *ChainNode) StartServer() {
	listen, err := net.Listen("tcp", n.Addr)
	if err != nil {
		fmt.Printf("node startup failed: %v\n", err)
		return
	}
	defer listen.Close()
	fmt.Printf("blockchain node running: %s\n", n.Addr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go n.handleConn(conn)
	}
}

func (n *ChainNode) handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	conn.Read(buf)
	n.Lock.Lock()
	n.Connected = append(n.Connected, conn.RemoteAddr().String())
	n.Lock.Unlock()
	fmt.Printf("new peer connected: %s\n", conn.RemoteAddr().String())
}

func main() {
	node := NewNode("0.0.0.0:9123")
	node.StartServer()
}
