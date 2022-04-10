package parser

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// handle functions
// where the parser found the KEYWORD_CONST code pattern

// base function for handling const expressions

// rules for variable names
//	* Must START with a character
//  * Must NOT have special characters (any type)
//	* Must ONLY have numbers AFTER the first CHARACTER

func handle_const_expression(cursor *multipleCursor, parser *Parser) {
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
	var v lang.Type
	if tokensAreString(cursor, lexer) {
		v = tokensToGagoString(cursor, lexer)
	} else if tokensAreInt(cursor, lexer) {
		fmt.Println("found int")
		v = tokensToGagoInt(cursor, lexer)
	}

	v.SetConstant(true)
	parser.Ast = append(parser.Ast, ast.VariableDeclaration{AstType: ast.AST_TYPE_VAR_DECLARATION, Vtype: ast.VTYPE_CONST, Vname: vname, Vvalue: v})
}

// TODO: put the below functions in a new utils directory
func tokensAreString(cursor *multipleCursor, lexer *lexer.Lexer) bool {
	qt := 0
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
				return false
			}
		} else if len(cursor.CurrentTokens)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return false
			}
		}
	}
	return true
}

func tokensToGagoString(cursor *multipleCursor, lexer *lexer.Lexer) *lang.TypeString {
	qt := 0
	tmpvalue := ""
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
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return nil
			}
		} else if len(cursor.CurrentTokens)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return nil
			}
		} else {
			tmpvalue += t.GetValue().(string)
		}
	}
	return lang.String(tmpvalue)
}

func tokensAreInt(cursor *multipleCursor, lexer *lexer.Lexer) bool {
	for i, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			if !(tkntest.NValueIs("-") || tkntest.IsNum()) {
				return false
			}
		} else {
			if !tkntest.IsNum() {
				return false
			}
		}
	}
	return true
}

func tokensToGagoInt(cursor *multipleCursor, lexer *lexer.Lexer) *lang.TypeInt {
	var str string
	for _, v := range cursor.CurrentTokens {
		q, ok := v.GetValue().(string)
		if !ok {
			panic("internal error: tokenstogagoint failed")
		}
		str += q
	}
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("internal error: tokenstogagoint failed")
	}
	return lang.Int(n)
}
