package main

import (
	"fmt"

	"github.com/alexey-belous/blockchain-lessons/core"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("Data 1")
	bc.AddBlock("Data 2")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()
	}
}
