package main

import (
	"math/big"
 
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	Auth *bind.TransactOpts
	Rpc *ethclient.Client
}

func (c *Client) SetNonce(nonce *big.Int) {
	c.Auth.Nonce = nonce
}

func (c *Client) SetFundValue(fundValue *big.Int) {
	c.Auth.Value = fundValue
}

func (c *Client) SetGasLimit(gasLimit uint64) {
	c.Auth.GasLimit = gasLimit
}

func (c *Client) SetGasPrice(gasPrice *big.Int) {
	c.Auth.GasPrice = gasPrice
}

func NewClient(endpoint string, privateKey string, chainId *big.Int) (Client, error) {
	rpc, err := ethclient.Dial(endpoint)
	if err != nil {
		return Client{}, errors.Wrap(err, "failed to make a new rpc client")
	}

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return Client{}, errors.Wrap(err, "invalid key: failed to convert private key from hex to ecdsa")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return Client{}, errors.Wrap(err, "invalid key: failed to create a transaction signer from private key")
	}

	return Client{auth, rpc}, nil
}
