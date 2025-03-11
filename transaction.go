package main

type TransactionInput struct {
	Hash       []byte
	Value      int
	Signature  []byte
	LockScript string
}

type TransactionOutput struct {
	Value        int
	UnlockScript string
}

type Transaction struct {
	Hash    []byte
	Inputs  []TransactionInput
	Outputs []TransactionOutput
	Nonce   int
}
