package ast

type VariableReDeclaration struct {
	AstType int         // ast type
	vname   string      // variable name
	vvalue  interface{} // new value
}
