package parser

import "github.com/glaukiol1/gago/ast"

// this file parses the
// import statements

func handle_import_expression(cursor *multipleCursor, parser *Parser) {
	cursor.SetIndex(1)
	var modname string
	for _, v := range cursor.CurrentTokens {
		modname += v.GetValue().(string)
	}
	parser.Ast = append(parser.Ast, ast.Import{AstType: ast.AST_TYPE_IMPORT, Mname: modname})
}
