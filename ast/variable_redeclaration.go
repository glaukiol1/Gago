package ast

type VariableReDeclaration struct {
	AstType int         // ast type
	Vname   string      // variable name
	Vvalue  interface{} // new value
}
