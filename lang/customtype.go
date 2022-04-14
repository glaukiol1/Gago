package lang

// custom type is a type where it would be specific to one function
// or created within the Gago script

// this type will be used for arrays, struct-like types and more

type CustomType struct {
	BaseType
}

func LoadCustomType(qualname string, value interface{}) *CustomType {
	return &CustomType{BaseType{Typename: qualname, Value: value}}
}
