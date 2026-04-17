package main

import (
	"fmt"
	"sync"
)

type UTXO struct {
	TxID  string
	Index int
	Addr  string
	Value float64
	Spent bool
}

type UTXOSet struct {
	Set  map[string]*UTXO
	Lock sync.RWMutex
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		Set: make(map[string]*UTXO),
	}
}

func (u *UTXOSet) Add(utxo *UTXO) {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	key := utxo.TxID + fmt.Sprintf("%d", utxo.Index)
	u.Set[key] = utxo
}

func (u *UTXOSet) Spend(txID string, index int) {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	key := txID + fmt.Sprintf("%d", index)
	if utxo, ok := u.Set[key]; ok {
		utxo.Spent = true
	}
}

func (u *UTXOSet) GetBalance(addr string) float64 {
	u.Lock.RLock()
	defer u.Lock.RUnlock()
	var total float64
	for _, utxo := range u.Set {
		if utxo.Addr == addr && !utxo.Spent {
			total += utxo.Value
		}
	}
	return total
}

func main() {
	set := NewUTXOSet()
	set.Add(&UTXO{TxID: "tx001", Index: 0, Addr: "addr1", Value: 50.0})
	fmt.Println("balance:", set.GetBalance("addr1"))
	set.Spend("tx001", 0)
	fmt.Println("balance after spend:", set.GetBalance("addr1"))
}
