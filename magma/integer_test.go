package magma

import (
	"github.com/cocoatomo/goarith/magma"
	"github.com/cocoatomo/goarith/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagmaIntegerOp(t *testing.T) {
	assert.Equal(t, types.NewInteger(2), magma.Op(types.NewInteger(1), types.NewInteger(1)))
}
