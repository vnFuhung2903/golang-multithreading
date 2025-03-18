package main

import (
	"crypto/sha256"
	"fmt"
	"gobtc/entities"
	"sync"
	"time"
)

const decimal int = 100000000

func main() {
	genesisHash := sha256.Sum256([]byte("VCS"))
	myBlockchain := entities.NewBlockchain(genesisHash)
	fmt.Println("Created blockchain")

	alice := entities.NewWallet()
	fmt.Println("Alice's wallet:", alice.Address())
	bob := entities.NewWallet()
	fmt.Println("Bob's wallet:", bob.Address())

	newBlock := myBlockchain.MineBlock(
		entities.NewCoinBaseTransaction(alice.Address(), 15*decimal),
		entities.NewCoinBaseTransaction(bob.Address(), 6*decimal),
	)
	myBlockchain.AddBlock(newBlock)

	aliceBalance, _ := myBlockchain.FindSpendableUTXO(alice.Address())
	fmt.Println("Alice's initial balance:", aliceBalance/decimal)
	bobBalance, _ := myBlockchain.FindSpendableUTXO(bob.Address())
	fmt.Println("Bob's initial balance:", bobBalance/decimal)

	var wg sync.WaitGroup
	var txs []*entities.Transaction
	var tx *entities.Transaction
	channel := make(chan string, 10)
	for i := 1; i < 6; i++ {
		wg.Add(1)
		tx = entities.NewTransaction(alice, bob.Address(), i*decimal, myBlockchain)
		go func(tx *entities.Transaction) {
			defer wg.Done()
			txs = append(txs, tx)
			channel <- fmt.Sprintf("Transfer from Alice to Bob %d btc", i)
			time.Sleep(1000 * time.Millisecond)
		}(tx)
	}
	for i := 1; i < 4; i++ {
		wg.Add(1)
		tx = entities.NewTransaction(bob, alice.Address(), i*decimal, myBlockchain)
		go func(tx *entities.Transaction) {
			defer wg.Done()
			txs = append(txs, tx)
			channel <- fmt.Sprintf("Transfer from Bob to Alice %d btc", i)
			time.Sleep(1000 * time.Millisecond)
		}(tx)
	}
	wg.Wait()
	close(channel)
	for str := range channel {
		fmt.Println(str)
	}

	newBlock = myBlockchain.MineBlock(txs...)
	myBlockchain.AddBlock(newBlock)

	aliceBalance, _ = myBlockchain.FindSpendableUTXO(alice.Address())
	fmt.Println("Alice's balance:", aliceBalance/decimal)
	bobBalance, _ = myBlockchain.FindSpendableUTXO(bob.Address())
	fmt.Println("Bob's balance:", bobBalance/decimal)
}
