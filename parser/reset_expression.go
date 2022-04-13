package parser

import (
	"github.com/glaukiol1/gago/ast"
)

// the `reset` expression
// to reset a variable
// example code:
/*
var test = "hello world"
call print(test)
reset test = "not hello world"
call print(test)
*/

func handle_reset_expression(cursor *multipleCursor, parser *Parser) {
	vname, vvalue, ok := vvhandler(cursor, parser)
	if !ok {
		return
	}
	parser.Ast = append(parser.Ast, ast.VariableReDeclaration{AstType: ast.AST_TYPE_VAR_REDECLARATION, Vname: vname, Vvalue: vvalue})
}
