package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Timestamp int64
	Signature string
}

func SignTx(tx *Transaction, priv *ecdsa.PrivateKey) error {
	hash := sha256.Sum256([]byte(tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)))
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return err
	}
	tx.Signature = hex.EncodeToString(r.Bytes()) + hex.EncodeToString(s.Bytes())
	return nil
}

func VerifyTx(tx *Transaction, pub *ecdsa.PublicKey) bool {
	hash := sha256.Sum256([]byte(tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)))
	sigLen := len(tx.Signature)
	rBytes, _ := hex.DecodeString(tx.Signature[:sigLen/2])
	sBytes, _ := hex.DecodeString(tx.Signature[sigLen/2:])
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)
	return ecdsa.Verify(pub, hash[:], r, s)
}

func main() {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tx := &Transaction{
		From:      "0x1a2b3c4d5e6f",
		To:        "0x9f8e7d6c5b4a",
		Amount:    10.25,
		Timestamp: 1735000000,
	}
	SignTx(tx, priv)
	fmt.Println("tx signed:", tx.Signature)
	fmt.Println("verify result:", VerifyTx(tx, &priv.PublicKey))
}
