package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type P2PMessage struct {
	Type    string
	Content string
	Sender  string
}

type Router struct {
	Port     string
	Handlers map[string]func(P2PMessage) string
}

func NewRouter(port string) *Router {
	return &Router{
		Port:     port,
		Handlers: make(map[string]func(P2PMessage) string),
	}
}

func (r *Router) Register(msgType string, handler func(P2PMessage) string) {
	r.Handlers[msgType] = handler
}

func (r *Router) Start() {
	listen, _ := net.Listen("tcp", ":"+r.Port)
	defer listen.Close()
	fmt.Println("p2p router running on port", r.Port)
	for {
		conn, _ := listen.Accept()
		go r.process(conn)
	}
}

func (r *Router) process(conn net.Conn) {
	defer conn.Close()
	var msg P2PMessage
	json.NewDecoder(conn).Decode(&msg)
	if fn, ok := r.Handlers[msg.Type]; ok {
		res := fn(msg)
		conn.Write([]byte(res))
	}
}

func main() {
	router := NewRouter("9001")
	router.Register("BLOCK_SYNC", func(msg P2PMessage) string {
		return "sync success"
	})
	router.Start()
}
