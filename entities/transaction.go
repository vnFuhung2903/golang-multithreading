package entities

import (
	"crypto/sha256"
)

type Transaction struct {
	Hash       [32]byte
	Inputs     []*TransactionInput
	Outputs    []*TransactionOutput
	Nonce      int
	IsCoinbase bool
}

type TransactionInput struct {
	Hash       [32]byte
	Value      int
	Signature  []byte
	LockScript string
}

type TransactionOutput struct {
	Value        int
	UnlockScript string
}

func NewCoinBaseTransaction(to string, amount int) *Transaction {
	TXO := &TransactionOutput{amount, to}
	return &Transaction{
		Inputs:     []*TransactionInput{},
		Outputs:    []*TransactionOutput{TXO},
		IsCoinbase: true,
		Nonce:      0,
		Hash:       sha256.Sum256([]byte(to)),
	}
}

func NewTransaction(from *Wallet, to string, amount int, blockchain *Blockchain) *Transaction {
	var TXIs []*TransactionInput
	var TXOs []*TransactionOutput
	fromAddress := from.Address()
	accumulated, spendableUTXO := blockchain.FindSpendableUTXO(fromAddress)
	if accumulated < amount {
		panic("ERROR: Not enough funds")
	}

	for _, value := range spendableUTXO {
		TXIs = append(TXIs, &TransactionInput{
			Hash:       [32]byte{},
			Signature:  []byte{0},
			Value:      value,
			LockScript: fromAddress,
		})
	}

	TXOs = append(TXOs, &TransactionOutput{
		Value:        amount,
		UnlockScript: to,
	})

	TXOs = append(TXOs, &TransactionOutput{
		Value:        accumulated - amount,
		UnlockScript: fromAddress,
	})

	tx := &Transaction{
		Inputs:     TXIs,
		Outputs:    TXOs,
		IsCoinbase: false,
		Nonce:      0,
		Hash:       [32]byte{},
	}
	return tx
}
