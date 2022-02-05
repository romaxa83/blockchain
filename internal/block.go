package internal

import (
	"bytes"
	"crypto/sha256"
)

// простая структура блока
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// создание блока
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// вычисление хеша
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
