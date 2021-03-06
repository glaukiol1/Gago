package ast

const VTYPE_VAR = 0
const VTYPE_CONST = 1

type VariableDeclaration struct {
	AstType int         // ast type
	Vtype   int         // constant or not constant
	Vname   string      // variable name
	Vvalue  interface{} // variable value
}
