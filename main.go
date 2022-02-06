package main

import (
	"flag"
	"fmt"
	"github.com/romaxa83/blockchain/internal"
	"os"
	"rsc.io/quote"
	"runtime"
	"strconv"
)

type CommandLine struct {
	blockchain *internal.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Nonce : %x\n", block.Nonce)

		pow := internal.NewProof(block)
		fmt.Printf("PoW : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		internal.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		internal.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	fmt.Println(quote.Hello())

	chain := internal.InitBlockChain()
	defer chain.DB.Close()

	cli := CommandLine{chain}
	cli.run()
}
