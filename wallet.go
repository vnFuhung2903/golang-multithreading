package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  [32]byte
}

func NewWallet() *Wallet {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := sha256.Sum256(append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...))
	return &Wallet{
		PrivateKey: *privateKey,
		PublicKey:  publicKey,
	}
}

func (wallet *Wallet) Address() string {
	hash := crypto.RIPEMD160.New()
	publicKeyHash, err := hash.Write(wallet.PublicKey[:])
	if err != nil {
		fmt.Println(err)
	}
	version := byte(0x00)
	check := checksum([]byte{byte(publicKeyHash)})

	address := append([]byte{version}, byte(publicKeyHash))
	address = append(address, check...)
	return base58.Encode(address)
}

func checksum(payload []byte) []byte {
	SHA1 := sha256.Sum256(payload)
	SHA2 := sha256.Sum256(SHA1[:])
	return SHA2[:4]
}
