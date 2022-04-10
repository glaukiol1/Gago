package lang

// this file holds the functions specific to
// the lang.String type
// find the definition in types.go

func String(s string) *TypeString {
	return &TypeString{BaseType{Typename: "string", Value: s, Constant: false}}
}

func (str *TypeString) Len() int {
	return len(str.Value.(string))
}
