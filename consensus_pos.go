package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Addr    string
	Stake   float64
	Active  bool
}

type PoS struct {
	Validators []*Validator
}

func NewPoS() *PoS {
	return &PoS{
		Validators: make([]*Validator, 0),
	}
}

func (p *PoS) Register(addr string, stake float64) {
	p.Validators = append(p.Validators, &Validator{
		Addr:   addr,
		Stake:  stake,
		Active: true,
	})
}

func (p *PoS) ElectBlockProducer() *Validator {
	rand.Seed(time.Now().UnixNano())
	total := 0.0
	for _, v := range p.Validators {
		if v.Active {
			total += v.Stake
		}
	}
	target := rand.Float64() * total
	current := 0.0
	for _, v := range p.Validators {
		if v.Active {
			current += v.Stake
			if current >= target {
				return v
			}
		}
	}
	return nil
}

func main() {
	pos := NewPoS()
	pos.Register("node1", 100)
	pos.Register("node2", 300)
	producer := pos.ElectBlockProducer()
	fmt.Println("producer:", producer.Addr)
}
