package main
	 
import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (c *Contract) MintToken(client Client, tokenId *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	tx, err := c.Instance.MintCaller(client.Auth, tokenId, tokenAmount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signed mint transaction")
	}

	return tx, nil
}
