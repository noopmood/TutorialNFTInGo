package main

import (
	"fmt"
	"math/big"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkOk(ok bool) {
	if !ok {
		panic(ok)
	}
}

func main() {
	endpoint := "${YOUR_ENDPOINT}"
	privateKey := "${YOUR_WALLET_PRIVATE_KEY}"
	chainId := big.NewInt(3)

	client, err := NewClient(endpoint, privateKey, chainId)
	check(err)

	client.SetNonce(${YOUR_WALLET_NONCE})
	client.SetFundValue(big.NewInt(0))
	client.SetGasLimit(uint64(8000000))
	client.SetGasPrice(big.NewInt(1875000000))

	contract, err := client.DeployContract()
	check(err)
	fmt.Println("Contract address: ", contract.Address.Hex())

	client.SetNonce(${YOUR_WALLET_NONCE})
	tokenId := new(big.Int)
	tokenId, ok := tokenId.SetString("0", 16)
	checkOk(ok)
	tokenAmount := new(big.Int)
	tokenAmount, ok = tokenAmount.SetString("5", 16)
	checkOk(ok)
	fmt.Println("Minting ", tokenAmount, " tokens of ID: ", tokenId)

	tx, err := contract.MintToken(client, tokenId, tokenAmount)
	check(err)
	fmt.Println("Mint transaction: ", tx.Hash())
}

