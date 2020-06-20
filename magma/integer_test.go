package magma

import (
	"github.com/cocoatomo/goarith/structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagmaIntegerOp(t *testing.T) {
	assert.Equal(t, structure.NewInteger(2), Op(structure.NewInteger(1), structure.NewInteger(1)))
}
