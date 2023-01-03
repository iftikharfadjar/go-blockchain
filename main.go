package main

import (
	"fmt"
	"block/block.go"
)

func main() {
	fmt.Println("hi")
	b := block.NewBlock(0, "init hash")
	b.Print()
}