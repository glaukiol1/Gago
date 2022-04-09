package ast

const VTYPE_VAR = 0
const VTYPE_CONST = 1

type VariableDeclaration struct {
	Ast                // inherit Ast struct
	Vtype  int         // constant or not constant
	Vname  string      // variable name
	Vvalue interface{} // variable value
}
