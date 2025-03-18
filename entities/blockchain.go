package entities

import (
	"errors"
	"time"
)

type Blockchain struct {
	blocks []*Block
}

func (blockchain *Blockchain) AddBlock(newBlock *Block) error {
	var err error
	if newBlock == nil || newBlock.Height < blockchain.blocks[len(blockchain.blocks)-1].Height {
		err = errors.New("ERROR: New block height is too low")
	}
	blockchain.blocks = append(blockchain.blocks, newBlock)
	return err
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

func (blockchain *Blockchain) MineBlock(txs ...*Transaction) *Block {
	for _, tx := range txs {
		if blockchain.checkTransactionExists(tx) {
			panic("ERROR: Transaction existed")
		}
	}
	newBlock := NewBlock(txs, blockchain.blocks[len(blockchain.blocks)-1])
	return newBlock
}

func (blockchain *Blockchain) FindSpendableUTXO(address string) (int, []int) {
	UTXOs := blockchain.findWalletTXO(address)
	var spendableUTXO []int
	res := 0

	for _, unspentTXO := range UTXOs {
		res += unspentTXO
		spendableUTXO = append(spendableUTXO, unspentTXO)
	}
	return res, spendableUTXO
}

func (blockchain *Blockchain) findWalletTXO(address string) []int {
	var UTXOs []int
	spentUTXOs := make(map[[32]byte][]int)

	for _, block := range blockchain.blocks {
		for _, tx := range block.Data {
			txSpent := spentUTXOs[tx.Hash]
			for outputId, output := range tx.Outputs {
				if checkTXOSpent(txSpent, outputId) {
					continue
				}
				if output.UnlockScript == address {
					UTXOs = append(UTXOs, output.Value)
				}
				if tx.IsCoinbase {
					continue
				}
			}
			for _, input := range tx.Inputs {
				if input.LockScript == address {
					UTXOs = append(UTXOs, -input.Value)
					spentUTXOs[input.Hash] = append(spentUTXOs[input.Hash], input.Value)
				}
			}
		}
	}
	return UTXOs
}

func checkTXOSpent(spentUTXOs []int, id int) bool {
	for _, spentUTXO := range spentUTXOs {
		if spentUTXO == id {
			return true
		}
	}
	return false
}

func (blockchain *Blockchain) checkTransactionExists(anonymousTx *Transaction) bool {
	for _, block := range blockchain.blocks {
		for _, tx := range block.Data {
			if tx.Hash == anonymousTx.Hash {
				return true
			}
		}
	}
	return false
}
