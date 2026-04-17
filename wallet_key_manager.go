package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := elliptic.Marshal(priv.Curve, priv.X, priv.Y)
	addr := hex.EncodeToString(pub[:20])
	return &Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
		Address:    addr,
	}
}

func (w *Wallet) ExportPriv() string {
	return hex.EncodeToString(w.PrivateKey.D.Bytes())
}

func main() {
	wallet := NewWallet()
	fmt.Println("address:", wallet.Address)
	fmt.Println("public key:", hex.EncodeToString(wallet.PublicKey))
}
