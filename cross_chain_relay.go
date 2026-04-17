package main

import (
	"fmt"
	"time"
)

type CrossChainMessage struct {
	SourceChain string
	TargetChain string
	Data        string
	Timestamp   int64
}

type Relay struct {
	Queue []*CrossChainMessage
}

func NewRelay() *Relay {
	return &Relay{
		Queue: make([]*CrossChainMessage, 0),
	}
}

func (r *Relay) Push(msg *CrossChainMessage) {
	msg.Timestamp = time.Now().Unix()
	r.Queue = append(r.Queue, msg)
	fmt.Println("relay message pushed")
}

func (r *Relay) Forward() {
	for len(r.Queue) > 0 {
		msg := r.Queue[0]
		r.Queue = r.Queue[1:]
		fmt.Printf("forward from %s to %s: %s\n", msg.SourceChain, msg.TargetChain, msg.Data)
	}
}

func main() {
	relay := NewRelay()
	relay.Push(&CrossChainMessage{
		SourceChain: "ETH",
		TargetChain: "BSC",
		Data:        "transfer 15",
	})
	relay.Forward()
}
