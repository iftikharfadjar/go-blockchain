package main

import (
	"fmt"

	"blockchain-go.com/block"
)

func main() {
	fmt.Println("hi")
	b := block.NewBlock(0, "init hash")
	b.Print()
}