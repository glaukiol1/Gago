package parser

import (
	"github.com/glaukiol1/gago/ast"
)

// handle functions
// where the parser found the KEYWORD_CONST code pattern

// base function for handling const expressions

// rules for variable names
//	* Must START with a character
//  * Must NOT have special characters (any type)
//	* Must ONLY have numbers AFTER the first CHARACTER

func handle_var_expression(cursor *multipleCursor, parser *Parser) {
	parser.Ast = append(parser.Ast, vhandler(cursor, parser, ast.VTYPE_VAR))
}
