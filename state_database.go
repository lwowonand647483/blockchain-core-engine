package main

import (
	"fmt"
	"sync"
)

type StateDB struct {
	Data map[string]string
	Lock sync.RWMutex
}

func NewStateDB() *StateDB {
	return &StateDB{
		Data: make(map[string]string),
	}
}

func (s *StateDB) Put(key, value string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Data[key] = value
}

func (s *StateDB) Get(key string) string {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return s.Data[key]
}

func (s *StateDB) Delete(key string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Data, key)
}

func main() {
	db := NewStateDB()
	db.Put("user1", "1000")
	fmt.Println(db.Get("user1"))
}
