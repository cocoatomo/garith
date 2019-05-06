package types

import "math/big"

type Integer = *big.Int

func NewInteger(x int64) Integer {
	return big.NewInt(x)
}
