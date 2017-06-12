package main

import (
	"fmt"
	//bk "github.com/blockchain-db/core/common"
	"time"
	"math/rand"
	_ "strconv"
	"crypto/sha256"
)

func main() {
	timestamp := time.Now().UnixNano()
	seed := fmt.Sprint(timestamp)

	rand.Seed(timestamp)
	fmt.Println(sha256.Sum256([]byte(seed)))
	//log.Print(fmt.Sprintf("%x", ))

	//block := bk.NewBlock()
	//fmt.Println(block)
}