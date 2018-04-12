package main

//структура для сборки блоков в единый блокчейн
type Blockchain struct {
	blocks []*Block
}

//добавление блоков в блокчейн
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//создание генезис блока
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//создаем новый блокчейн
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
