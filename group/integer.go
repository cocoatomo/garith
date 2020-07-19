package group

import (
	"math/big"
)

type Integer struct {
	value *big.Int
}

func NewInteger(value int64) Integer {
	return Integer{big.NewInt(value)}
}

func (i Integer) Op(i2 Integer) Integer {
	result := NewInteger(0)
	return Integer{result.value.Add(i.value, i2.value)}
}

func (i Integer) Identity() Integer {
	return NewInteger(0)
}

func (i Integer) Inverse() Integer {
	inverse := NewInteger(0)
	return Integer{inverse.value.Neg(i.value)}
}
