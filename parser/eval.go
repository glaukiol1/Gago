package parser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/glaukiol1/gago/ast"
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
	var chars string
	for _, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, parser.lexer)
		codes = append(codes, tkntest.token.GetCode())
		chars += t.GetValue().(string)
		if reflect.DeepEqual(codes, KEYWORD_CALL_CODE) {
			if parser.v {
				fmt.Println("found call statement assignment")
			}
			return nhandle_call_expression(cursor, parser, false)
		}
	}
	if tokensAreString(cursor, parser.lexer) {
		return tokensToGagoString(cursor, parser.lexer)
	} else if tokensAreInt(cursor, parser.lexer) {
		return tokensToGagoInt(cursor, parser.lexer)
	} else {
		if exprIsMathEquation(chars) {
			if parser.v {
				fmt.Println("found math expression")
			}
			return evalMathExpr(chars)
		} else {
			return ast.VariableAccess{AstType: ast.AST_TYPE_VARIABLE_ACCESS, Vname: chars}
		}
	}
	// lang.Errorf("SyntaxError", "Unknown expression type.", lang.BuildStack(cursor.CurrentTokens[len(cursor.CurrentTokens)-1], parser.lexer.GetFilename()), true).Run()
	// return nil
}

// evaluate if expr is math expression
func exprIsMathEquation(s string) bool {
	if strings.ContainsAny(s, "+-^*/") {
		_, err := govaluate.NewEvaluableExpression(s)
		return err == nil
	}
	return false
}

// eval math equation
func evalMathExpr(s string) *ast.MathExpr {
	if !exprIsMathEquation(s) {
		panic("internal error: evalMathExpr failed")
	}
	_, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		panic("internal error: evalMathExpr failed")
	}
	return &ast.MathExpr{AstType: ast.AST_TYPE_MATH_EXPR, Expression: s}
}
