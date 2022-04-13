package lang

// the two boolean values for Gago

var False = &TypeBool{BaseType{Typename: "bool", Value: false, Constant: true}}
var True = &TypeBool{BaseType{Typename: "bool", Value: true, Constant: true}}
