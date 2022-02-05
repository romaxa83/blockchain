package main

import (
	"fmt"
	"github.com/romaxa83/blockchain/internal"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	chain := internal.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %x\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
	}
}
