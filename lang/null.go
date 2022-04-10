package lang

type null struct {
	BaseType
}

var Null = null{BaseType: BaseType{Typename: "null", Value: nil, Constant: true}}
