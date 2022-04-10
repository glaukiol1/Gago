package parser

import "github.com/glaukiol1/gago/ast"

// this is only for testing...

// unknown expression will get a variable name that is the only
// text in that line, and set out a VariableAccess AST for the VM
// to proccess it and show how variables can be accessed.

func handle_unknown_expression(cursor *multipleCursor, parser *Parser) {

	cursor.SetIndex(0) // unknown expression will only have one part to them

	vname := ""
	for _, t := range cursor.CurrentTokens {
		vname += t.GetValue().(string)
	}

	_ast := ast.VariableAccess{AstType: ast.AST_TYPE_VARIABLE_ACCESS, Vname: vname}
	parser.Ast = append(parser.Ast, _ast)
}
