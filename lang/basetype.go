package lang

// base type

// all the other types derive from this type

type BaseType struct {
	Typename string
	Value    interface{}
}

func (typ *BaseType) Reassign(newval interface{}) {
	typ.Value = newval
}

func (typ *BaseType) Val() interface{} {
	return typ.Value
}

func (typ *BaseType) Name() string {
	return typ.Typename
}
