package structure

type ZFC interface {
     Contain(ZFC) bool
}

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
	Op(CommutativeGroup) CommutativeGroup
	Identity() CommutativeGroup
	Inverse(CommutativeGroup) CommutativeGroup
}

type AdditiveGroup interface {
	Op(AdditiveGroup) AdditiveGroup
	Identity() AdditiveGroup
	Inverse(AdditiveGroup) AdditiveGroup
}

type MultiplicativeGroup interface {
	Op(MultiplicativeGroup) MultiplicativeGroup
	Identity() MultiplicativeGroup
	Inverse(MultiplicativeGroup) MultiplicativeGroup
}

type Ring interface {
	Add(Ring) Ring
	Minus(Ring) Ring
	Multiply(Ring) Ring
	Zero() Ring
	One() Ring
}

type Field interface {
	Add(Field) Field
	Minus(Field) Field
	Multiply(Field) Field
	divide(Field) Field
	Zero() Field
	One() Field
}