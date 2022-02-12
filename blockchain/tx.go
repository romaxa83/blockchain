package blockchain

import (
	"bytes"
	"github.com/romaxa83/blockchain/wallet"
)

// выход, хранит данные (монеты)
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// вход, ссылается на предыдущий выход
type TxInput struct {
	ID        []byte // идентификатор транзакции выхода
	Out       int    // хранит индекс выхода данной транзакции
	Signature []byte
	PubKey    []byte
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
