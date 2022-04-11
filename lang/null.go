package lang

type NullType struct {
	BaseType
}

var Null = &NullType{BaseType: BaseType{Typename: "null", Value: nil, Constant: true}}
