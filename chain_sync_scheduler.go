package main

import (
	"fmt"
	"time"
)

type SyncTask struct {
	PeerAddr string
	Height   int
}

type Scheduler struct {
	Tasks []*SyncTask
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Tasks: make([]*SyncTask, 0),
	}
}

func (s *Scheduler) AddTask(addr string, height int) {
	s.Tasks = append(s.Tasks, &SyncTask{
		PeerAddr: addr,
		Height:   height,
	})
}

func (s *Scheduler) Run() {
	for _, task := range s.Tasks {
		fmt.Printf("sync height %d from %s\n", task.Height, task.PeerAddr)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	sch := NewScheduler()
	sch.AddTask("192.168.1.100", 500)
	sch.AddTask("192.168.1.101", 500)
	sch.Run()
}
