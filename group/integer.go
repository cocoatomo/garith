package group

type Integer int64

func (lhs Integer) Op(rhs Integer) Integer {
	return lhs + rhs
}

func (_ Integer) Identity() Integer {
	return 0
}

func (i Integer) Inverse() Integer {
	return - i
}
