package main

import (
	"fmt"

	"github.com/XhinoKurtaj/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after geneses")
	chain.AddBlock("Second block after geneses")
	chain.AddBlock("Third block after geneses")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
