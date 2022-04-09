package parser

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// handle functions
// where the parser found the KEYWORD_CONST code pattern

// some codes for certain keywords, expressions
const keyword_const = 0
const keyword_var = 1

// base function for handling const expressions

// rules for variable names
//	* Must START with a character
//  * Must NOT have special characters (any type)
//	* Must ONLY have numbers AFTER the first CHARACTER

func handle_const_expression(cursor *multipleCursor, lexer *lexer.Lexer) {
	cursor.SetIndex(1) // start at index 1

	// check if variable name is valid
	for i, tkn := range cursor.CurrentTokens {
		tkntest := NewTokenTest(tkn, lexer)
		if i == 0 {
			tkntest.IsChar(lang.Errorf("SyntaxError", "Expected character in variable name", lang.BuildStack(tkn, lexer.GetFilename()), true))
		} else {
			tkntest.IsNotSpecial(true)
		}
	}

	cursor.SetIndex(2) // switch to the index where the `=` should be located

	// checks for `=`
	if len(cursor.CurrentTokens) != 1 {
		lang.Errorf("SyntaxError", "Unexpected indentifier, expected `=`", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
		return
	}
	NewTokenTest(cursor.CurrentTokens[0], lexer).ValueIs("=", true)
}
