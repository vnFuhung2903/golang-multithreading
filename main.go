package main

import (
	"crypto/sha256"
	"fmt"
	"gobtc/entities"
	"gobtc/multithreading"
)

const decimal int = 100000000

func main() {
	multithreading.FetchURL()

	genesisHash := sha256.Sum256([]byte("VCS"))
	myBlockchain := entities.NewBlockchain(genesisHash)
	fmt.Println("Created blockchain:", myBlockchain)

	alice := entities.NewWallet()
	fmt.Println("Alice's wallet:", alice.Address())
	bob := entities.NewWallet()
	fmt.Println("Bob's wallet:", bob.Address())

	tx := entities.NewCoinBaseTransaction(alice.Address(), 20*decimal)
	newBlock := myBlockchain.MineBlock(tx)
	err := myBlockchain.AddBlock(newBlock)
	if err != nil {
		panic(err)
	}
	aliceBalance, _ := myBlockchain.FindSpendableUTXO(alice.Address())
	fmt.Println("Alice's balance:", aliceBalance/decimal)

	tx = entities.NewTransaction(alice, bob.Address(), 1*decimal, myBlockchain)
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
