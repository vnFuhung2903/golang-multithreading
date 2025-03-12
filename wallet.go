package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
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
	publicKeyHash, err := wallet.HashPublicKey()
	if err != nil {
		fmt.Println(err)
	}
	version := byte(0x00)
	check := checksum(publicKeyHash)

	address := append([]byte{version}, publicKeyHash...)
	address = append(address, check...)
	return base58.Encode(address)
}

func (wallet *Wallet) HashPublicKey() ([]byte, error) {
	hash := ripemd160.New()
	_, err := hash.Write(wallet.PublicKey[:])
	return hash.Sum(wallet.PublicKey[:]), err
}

func checksum(payload []byte) []byte {
	SHA1 := sha256.Sum256(payload)
	SHA2 := sha256.Sum256(SHA1[:])
	return SHA2[:4]
}
