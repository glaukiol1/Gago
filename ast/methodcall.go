package ast

// calling a function

type FuncCall struct {
	AstType  int           // the ast type
	Funcname string        // the function name
	Args     []interface{} // the interface{} will be of type ast.VariableAccess or ast.Literal
}
