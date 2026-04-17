package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Peer struct {
	IP    string
	Port  int
	Score int
}

type Discovery struct {
	Peers []*Peer
}

func NewDiscovery() *Discovery {
	return &Discovery{
		Peers: make([]*Peer, 0),
	}
}

func (d *Discovery) AddPeer(ip string, port int) {
	rand.Seed(time.Now().UnixNano())
	d.Peers = append(d.Peers, &Peer{
		IP:    ip,
		Port:  port,
		Score: rand.Intn(100),
	})
}

func (d *Discovery) BestPeer() *Peer {
	if len(d.Peers) == 0 {
		return nil
	}
	best := d.Peers[0]
	for _, p := range d.Peers {
		if p.Score > best.Score {
			best = p
		}
	}
	return best
}

func main() {
	dis := NewDiscovery()
	dis.AddPeer("1.1.1.1", 9000)
	dis.AddPeer("2.2.2.2", 9000)
	best := dis.BestPeer()
	fmt.Println("best peer:", best.IP)
}
