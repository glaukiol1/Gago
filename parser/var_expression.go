package parser

import (
	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
)

// handle functions
// where the parser found the KEYWORD_VAR code pattern

// base function for handling var expressions

// rules for variable names
//	* Must START with a character
//  * Must NOT have special characters (any type)
//	* Must ONLY have numbers AFTER the first CHARACTER

func handle_var_expression(cursor *multipleCursor, parser *Parser) {
	lexer := parser.lexer
	cursor.SetIndex(1) // start at index 1

	// check if variable name is valid
	vname := ""
	for i, tkn := range cursor.CurrentTokens {
		tkntest := NewTokenTest(tkn, lexer)
		if i == 0 {
			tkntest.IsChar(lang.Errorf("SyntaxError", "Expected character in variable name", lang.BuildStack(tkn, lexer.GetFilename()), true))
		} else {
			tkntest.IsNotSpecial(true)
		}
		vname += tkn.GetValue().(string)
	}

	cursor.SetIndex(2) // switch to the index where the `=` should be located

	// checks for `=`
	if len(cursor.CurrentTokens) != 1 {
		lang.Errorf("SyntaxError", "Unexpected indentifier, expected `=`", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
		return
	}
	NewTokenTest(cursor.CurrentTokens[0], lexer).ValueIs("=", true)

	cursor.SetIndex(3) // switch to the value of the variable

	// checks for variable value
	qt := -1       // quote type, 0 for single 1 for double
	tmpvalue := "" // hold the string value for now
	//TODO: support more than just strings
	for i, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if isSq {
				qt = 0
			} else if isDq {
				qt = 1
			} else {
				lang.Errorf("SyntaxError", "Unexpected indentifier, expected a `'` or `\"`", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return
			}
		} else if len(cursor.CurrentTokens)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return
			}
		} else {
			tmpvalue += t.GetValue().(string)
		}
	}
	parser.Ast = append(parser.Ast, ast.VariableDeclaration{ast.AST_TYPE_VAR_DECLARATION, ast.VTYPE_VAR, vname, lang.String(tmpvalue)})
}
