package lang

// this file holds the functions specific to
// the lang.Float type
// find the definition in types.go

func Float(s float64) *TypeFloat {
	return &TypeFloat{BaseType{Typename: "float", Value: s, Constant: false}}
}
