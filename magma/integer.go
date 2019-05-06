package magma

import (
	"github.com/cocoatomo/goarith/types"
)

func Op(left, right types.Integer) types.Integer {
	var result = types.NewInteger(0)
	return result.Add(left, right)
}
