package ast

import "github.com/glaukiol1/gago/lang"

type VariableReDeclaration struct {
	AstType int       // ast type
	Vname   string    // variable name
	Vvalue  lang.Type // new value
}
