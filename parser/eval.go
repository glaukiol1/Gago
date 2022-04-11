package parser

import (
	"fmt"
	"reflect"

	"github.com/glaukiol1/gago/lang"
)

// this file will include functions to evaluate a expression
// for example, creating a variable, and the
// variable declaration looks something like:
/*
const name = call input("Enter your name: ")
*/

// evalexpr will evaluate an expression and return the AST
// for that expression
func evaltokens(cursor *multipleCursor, parser *Parser) interface{} {
	cursor = cursor.SubCursor(cursor.currentIndex)
	var codes []int
	for _, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, parser.lexer)
		codes = append(codes, tkntest.token.GetCode())
		if reflect.DeepEqual(codes, KEYWORD_CALL_CODE) {
			if parser.v {
				fmt.Println("found call statement assignment")
			}
			return nhandle_call_expression(cursor, parser)
		}
	}
	lang.Errorf("SyntaxError", "Unknown expression type.", lang.BuildStack(cursor.CurrentTokens[len(cursor.CurrentTokens)-1], parser.lexer.GetFilename()), true).Run()
	return nil
}
