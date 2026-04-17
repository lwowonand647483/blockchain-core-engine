package main

import (
	"fmt"
	"time"
)

type Heartbeat struct {
	NodeID  string
	Interval time.Duration
	Stop    chan bool
}

func NewHeartbeat(nodeID string, interval time.Duration) *Heartbeat {
	return &Heartbeat{
		NodeID:   nodeID,
		Interval: interval,
		Stop:     make(chan bool),
	}
}

func (h *Heartbeat) Start() {
	ticker := time.NewTicker(h.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("node %s alive\n", h.NodeID)
		case <-h.Stop:
			return
		}
	}
}

func main() {
	hb := NewHeartbeat("node-01", 1*time.Second)
	go hb.Start()
	time.Sleep(3 * time.Second)
	hb.Stop <- true
}
