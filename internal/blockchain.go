package internal

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

const dbPath = "./tmp/blocks"

type BlockChain struct {
	LastHash []byte
	DB       *badger.DB
}

// итератор по блокчейну
type BlockChainIterator struct {
	CurrentHash []byte
	DB          *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	// настраиваем/открываем соединение к бд
	//opts := badger.DefaultOptions(dbPath) v1.6

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath
	db, err := badger.Open(opts)
	Handle(err)

	// "lh" - lastHash
	err = db.Update(func(txn *badger.Txn) error {
		// если нет записей, создаем первый блок
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()

			return err
		}
	})
	Handle(err)

	blockchain := BlockChain{lastHash, db}

	return &blockchain
}

// добавление блока в блокчейн
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})

	Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.DB}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		encodedBlock, err := item.Value()
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
