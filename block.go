package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// Block contains its data, hash, timestamp and previous block hash
type Block struct {
	Timestamp int64
	Hash      [32]byte
	Data      []*Transaction
	PrevHash  [32]byte
	Height    int
}

func NewBlock(txs []*Transaction, prevBlock *Block) *Block {
	timestamp := time.Now().Unix()
	blockData := HashAllTransactions(txs)
	blockHeader := bytes.Join([][]byte{{byte(timestamp)}, prevBlock.Hash[:], blockData[:]}, []byte{})
	blockHash := sha256.Sum256(blockHeader)
	block := &Block{
		Timestamp: timestamp,
		Data:      txs,
		PrevHash:  prevBlock.Hash,
		Hash:      blockHash,
		Height:    prevBlock.Height + 1,
	}
	return block
}

func HashAllTransactions(txs []*Transaction) [32]byte {
	var blockData []byte
	for _, tx := range txs {
		blockData = append(blockData, tx.Hash[:]...)
	}
	return sha256.Sum256(blockData)
}
