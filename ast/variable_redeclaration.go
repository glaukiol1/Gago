package ast

type VariableReDeclaration struct {
	vname  string      // variable name
	vvalue interface{} // new value
	Ast                // inherit Ast struct
}
