package lang

// this file holds all default types
// of gago programming language

type TypeString struct {
	BaseType
}

func String(s string) *TypeString {
	return &TypeString{BaseType{Typename: "string", Value: s}}
}
