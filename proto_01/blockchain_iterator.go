package main

import (
	"log"

	"bolt"
)

// структура для обхода блоков в базе
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// возвращаем следующий(тут значит предыдущий по блокчейну) блок
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
