package internal

// простая структура блока
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// создание блока
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
		0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
