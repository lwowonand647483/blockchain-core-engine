package main

import (
	"fmt"
	"time"
)

type Dashboard struct {
	NodeCount    int
	BlockHeight  int
	TxPerSecond  float64
	LastUpdateAt int64
}

func (d *Dashboard) Refresh(height int, count int, tps float64) {
	d.BlockHeight = height
	d.NodeCount = count
	d.TxPerSecond = tps
	d.LastUpdateAt = time.Now().Unix()
}

func (d *Dashboard) Show() {
	fmt.Printf("=== Chain Monitor ===\n")
	fmt.Printf("Height: %d\n", d.BlockHeight)
	fmt.Printf("Nodes: %d\n", d.NodeCount)
	fmt.Printf("TPS: %.2f\n", d.TxPerSecond)
	fmt.Printf("Update Time: %d\n", d.LastUpdateAt)
}

func main() {
	dash := Dashboard{}
	dash.Refresh(1000, 8, 12.5)
	dash.Show()
}
