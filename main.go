package main

import (
	"fmt"

	"github.com/BLoHny/Go_coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()

	fmt.Println(chain)
}
