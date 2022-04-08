package ast

const VTYPE_VAR = 0
const VTYPE_CONST = 1

type VariableDeclaration struct {
	vtype  int         // constant or not constant
	vname  string      // variable name
	vvalue interface{} // variable value
}
