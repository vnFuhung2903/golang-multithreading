package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	genesisHash := sha256.Sum256([]byte("VCS"))
	myBlockchain := NewBlockchain(genesisHash)
	fmt.Println(myBlockchain)

	wallet := NewWallet()
	fmt.Println(wallet)
}
