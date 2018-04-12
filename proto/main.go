package main

import "fmt"

//главная функция

func main() {
	bc := NewBlockchain() // создает блокчейн с генезис блоком
	// создает 2 блока в блокчейне
	fmt.Println("Add new block")
	bc.AddBlock("Send 1 BTC to Ivan")
	fmt.Println("Add new block")
	bc.AddBlock("Send 2 more BTC to Ivan")
	fmt.Println()

	//выводит все блоки блокчейна
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
