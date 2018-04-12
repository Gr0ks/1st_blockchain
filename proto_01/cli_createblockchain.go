package main

import (
	"fmt"
	"log"
)

// создаем блокчейн
func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
