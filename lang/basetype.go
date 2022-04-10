package lang

// base type

// all the other types derive from this type

type BaseType struct {
	Typename string
	Value    interface{}
	Constant bool // is the value is constant
}

func (typ *BaseType) Reassign(newval interface{}) error {
	if typ.Constant {
		return Errorf("TypeError", "Assignment to constant variable `"+typ.Name()+"`", "", true)
	}
	typ.Value = newval
	return nil
}

func (typ *BaseType) Val() interface{} {
	return typ.Value
}

func (typ *BaseType) Name() string {
	return typ.Typename
}

func (typ *BaseType) IsConstant() bool {
	return typ.Constant
}

func (typ *BaseType) SetConstant(to bool) {
	typ.Constant = to
}
