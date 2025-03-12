package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	genesisHash := sha256.Sum256([]byte("VCS"))
	myBlockchain := NewBlockchain(genesisHash)
	fmt.Println(myBlockchain)

	alice := NewWallet()
	fmt.Println("Alice's wallet: ", alice.Address())
	bob := NewWallet()
	fmt.Println("Bob's wallet: ", bob.Address())
}
