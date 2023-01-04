package block

import (
	"log"
	"time"
	"fmt"
)

type Block struct {
	nonce int
	previouseHash string
	timestamp int64
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("timestamp 		%d\n", b.timestamp)
	fmt.Printf("nonce 			%d\n", b.nonce)
	fmt.Printf("previous_hash 	%s\n", b.previouseHash)
	fmt.Printf("transaction 	%s\n", b.transactions)
}

func NewBlock(nonce int, previouseHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previouseHash = previouseHash
	return b
	
	// return &Block {
	// 	timestamp : time.Now().UnixNano(),
	// }
}


func init(){
	log.SetPrefix("Blockchain: ")
}

