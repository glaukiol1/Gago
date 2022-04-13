package lang

// this file holds the data specific to
// the lang.IntType type
// find the definition in types.go

func Int(s int64) *TypeInt {
	return &TypeInt{BaseType{Typename: "int", Value: s, Constant: false}}
}
