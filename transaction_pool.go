package main

import (
	"fmt"
	"sync"
)

type Tx struct {
	ID     string
	From   string
	To     string
	Amount float64
}

type TxPool struct {
	Pool map[string]*Tx
	Lock sync.Mutex
}

func NewTxPool() *TxPool {
	return &TxPool{
		Pool: make(map[string]*Tx),
	}
}

func (t *TxPool) Add(tx *Tx) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.Pool[tx.ID] = tx
}

func (t *TxPool) Remove(id string) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	delete(t.Pool, id)
}

func (t *TxPool) List() []*Tx {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	list := make([]*Tx, 0, len(t.Pool))
	for _, tx := range t.Pool {
		list = append(list, tx)
	}
	return list
}

func main() {
	pool := NewTxPool()
	pool.Add(&Tx{ID: "tx1", From: "a", To: "b", Amount: 5})
	fmt.Println("pool size:", len(pool.List()))
}
