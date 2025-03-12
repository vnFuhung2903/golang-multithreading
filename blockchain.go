package main

import "time"

type Blockchain struct {
	blocks []*Block
}

func (blockchain *Blockchain) AddBlock(data []*Transaction) {
	prevBlock := blockchain.blocks[len(blockchain.blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	blockchain.blocks = append(blockchain.blocks, newBlock)
}

func NewBlockchain(genesisHash [32]byte) *Blockchain {
	genesisBlock := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []*Transaction{},
		PrevHash:  [32]byte{},
		Hash:      genesisHash,
		Height:    0,
	}
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}

func (blockchain *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var UTXs []Transaction
	spentUTXOs := make(map[[32]byte][]int)

	for _, block := range blockchain.blocks {
		for _, tx := range block.Data {
			txSpent := spentUTXOs[tx.Hash]
			for outputId, output := range tx.Outputs {
				if checkTXOSpent(txSpent, outputId) {
					continue
				}
				if output.UnlockScript == address {
					UTXs = append(UTXs, *tx)
				}
				if tx.IsCoinbase {
					continue
				}
				for _, input := range tx.Inputs {
					if input.LockScript == address {
						spentUTXOs[input.Hash] = append(spentUTXOs[input.Hash], input.Value)
					}
				}
			}
		}
	}
	return UTXs
}

func (blockchain *Blockchain) FindSpendableUTXO(address string) (int, []int) {
	unspentTx := blockchain.FindUnspentTransactions(address)
	var spendableUTXO []int
	res := 0

	for _, tx := range unspentTx {
		for _, txo := range tx.Outputs {
			if txo.UnlockScript == address {
				res += txo.Value
				spendableUTXO = append(spendableUTXO, txo.Value)
			}
		}
	}
	return res, spendableUTXO
}

func checkTXOSpent(spentUTXOs []int, id int) bool {
	for _, spentUTXO := range spentUTXOs {
		if spentUTXO == id {
			return true
		}
	}
	return false
}
