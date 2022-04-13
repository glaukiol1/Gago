package parser

import (
	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
)

// handle variable declarations
// merge var_expression & const_expression
// so it isnt needed to copy the file contents

func vvhandler(cursor *multipleCursor, parser *Parser) (string, interface{}, bool) {
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
		return "", nil, false // will never run
	}
	NewTokenTest(cursor.CurrentTokens[0], lexer).ValueIs("=", true)

	cursor.SetIndex(3) // switch to the value of the variable

	// checks for variable value
	var v interface{}
	if s, ok := tokensString(cursor, parser.lexer); ok {
		v = s
	} else if tokensAreInt(cursor, lexer) {
		v = tokensToGagoInt(cursor, lexer)
	} else if tokensAreFloat(cursor, lexer) {
		v = tokensToFloat(cursor, lexer)
	} else {
		// check for expression type
		cursor.JoinAllFrom(3, " ")
		v = evaltokens(cursor, parser)
		if v == nil {
			return "", nil, false
		}
	}
	return vname, v, true
}

func vhandler(cursor *multipleCursor, parser *Parser, vtype int) (ast.VariableDeclaration, bool) {
	vname, vvalue, ok := vvhandler(cursor, parser)
	if !ok {
		return ast.VariableDeclaration{}, false
	}
	return ast.VariableDeclaration{AstType: ast.AST_TYPE_VAR_DECLARATION, Vtype: vtype, Vname: vname, Vvalue: vvalue}, true
}
