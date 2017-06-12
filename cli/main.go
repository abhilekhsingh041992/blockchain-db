package main

import (
	"fmt"
	bk "github.com/blockchain-db/core/common"
)

func main() {
	block := bk.NewBlock()
	fmt.Println(block)
}