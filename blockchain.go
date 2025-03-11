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
		Data:      [32]byte{},
		PrevHash:  [32]byte{},
		Hash:      genesisHash,
		Height:    0,
	}
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}
