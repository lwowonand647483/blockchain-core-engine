package main

import (
	"fmt"
	"sync"
)

type BlockPool struct {
	Blocks map[int]string
	Lock   sync.Mutex
}

func NewBlockPool() *BlockPool {
	return &BlockPool{
		Blocks: make(map[int]string),
	}
}

func (b *BlockPool) Add(height int, hash string) {
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.Blocks[height] = hash
}

func (b *BlockPool) Remove(height int) {
	b.Lock.Lock()
	defer b.Lock.Unlock()
	delete(b.Blocks, height)
}

func (b *BlockPool) Get(height int) (string, bool) {
	b.Lock.Lock()
	defer b.Lock.Unlock()
	h, ok := b.Blocks[height]
	return h, ok
}

func main() {
	pool := NewBlockPool()
	pool.Add(100, "0xabc123")
	hash, ok := pool.Get(100)
	fmt.Println(hash, ok)
}
