package main

import (
	"crypto/sha256"
	"fmt"
)

const decimal = 100000000

func main() {
	genesisHash := sha256.Sum256([]byte("VCS"))
	myBlockchain := NewBlockchain(genesisHash)
	fmt.Println(myBlockchain)

	alice := NewWallet()
	fmt.Println("Alice's wallet:", alice.Address())
	bob := NewWallet()
	fmt.Println("Bob's wallet:", bob.Address())

	tx := NewCoinBaseTransaction(alice.Address(), 20*decimal)
	newBlock := myBlockchain.MineBlock(tx)
	err := myBlockchain.AddBlock(newBlock)
	if err != nil {
		panic(err)
	}
	aliceBalance, _ := myBlockchain.FindSpendableUTXO(alice.Address())
	fmt.Println("Alice's balance:", aliceBalance/decimal)

	tx = NewTransaction(alice, bob.Address(), 1*decimal, myBlockchain)
	newBlock = myBlockchain.MineBlock(tx)
	err = myBlockchain.AddBlock(newBlock)
	if err != nil {
		panic(err)
	}
	aliceBalance, _ = myBlockchain.FindSpendableUTXO(alice.Address())
	fmt.Println("Alice's balance:", aliceBalance/decimal)
	bobBalance, _ := myBlockchain.FindSpendableUTXO(bob.Address())
	fmt.Println("Alice's balance:", bobBalance/decimal)
}
