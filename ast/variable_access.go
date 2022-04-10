package ast

// this is AST for when accessing a variable is needed

type VariableAccess struct {
	AstType int    // the ast type
	Vname   string // the variable name
}
