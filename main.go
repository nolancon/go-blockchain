package main

import (
	"fmt"
	"strconv"

	"github.com/nolancon/go-blockchain/pkg/blockchain"
)

func main() {
	bc := blockchain.InitBlockchain("genesis block!")

	bc.AddBlock("Second block!")
	bc.AddBlock("Third block!")
	bc.AddBlock("Fourth block!")

	for _, block := range bc.Blocks {
		fmt.Printf("Previous Block Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Nonce: %v\n", block.Nonce)
		fmt.Println()
	}
}
