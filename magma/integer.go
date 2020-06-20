package magma

type Integer int64

func (lhs Integer) Op(rhs Integer) Integer {
	return lhs * rhs
}
