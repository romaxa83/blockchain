package main

import (
	"fmt"
	"github.com/romaxa83/blockchain/internal"
	"rsc.io/quote"
	"strconv"
)

func main() {
	fmt.Println(quote.Hello())

	chain := internal.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Nonce : %x\n", block.Nonce)

		pow := internal.NewProof(block)
		fmt.Printf("PoW : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
