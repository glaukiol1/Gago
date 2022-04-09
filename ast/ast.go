package ast

const AST_TYPE_VAR_DECLARATION = 0
const AST_TYPE_VAR_REDECLARATION = 1

type Ast struct {
	asttype int // ast type
}
