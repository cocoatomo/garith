package structure

type Magma interface {
	Op(Magma) Magma
}

type Monoid interface {
	Op(Monoid) Monoid
	Identity() Monoid
}

type Semigroup interface {
	Op(Semigroup) Semigroup
}

type Group interface {
	Op(Group) Group
	Identity() Group
	Inverse(Group) Group
}

type CommutativeGroup interface {
	Op(Group) Group
	Identity() Group
	Inverse(Group) Group
}

type Ring interface {
	Add(Ring) Ring
	Multiply(Ring) Ring
	Zero() Ring
	One() Ring
}
