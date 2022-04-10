package parser

import (
	"fmt"
	"strings"

	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
	"github.com/glaukiol1/gago/utils"
)

// handle the call ... expression
// example expression:
// call print(hello)

// will subsitute variables inside the () with VariableAccess ast
// if its a literal, subsitute it with a newly created ast.Literal

func handle_call_expression(cursor *multipleCursor, parser *Parser) {
	tkns := cursor.JoinAllFrom(1) // join all tokens since the `call` keyword
	ok := false
	idx := 0
	funcname := ""
	for i, t := range tkns {
		tkntest := NewTokenTest(t, parser.lexer)

		f := tkntest.NValueIs("(")

		if !f {
			funcname += t.GetValue().(string)
		} else {
			ok = true
			idx = i
			break
		}
	}
	if !ok {
		lang.Errorf("SyntaxError", "Expected (", "", true).Run()
	}
	if parser.v {
		fmt.Println("Funcname: |" + funcname + "|")
	}

	tkns = tkns[idx:]

	rawargs := ""
	for i, t := range tkns {
		tkntest := NewTokenTest(t, parser.lexer)
		if parser.v {
			fmt.Println("Token: " + t.GetValue().(string))
		}
		if i == 0 {
			tkntest.ValueIs("(", true)
		} else if i == len(tkns)-1 {
			tkntest.ValueIs(")", true)
		} else {
			rawargs += t.GetValue().(string)
		}
	}
	var args []interface{}
	if len(rawargs) == 0 {
		if parser.v {
			fmt.Println("function no args")
		}
		args = make([]interface{}, 0)
	} else {
		if parser.v {
			fmt.Println("rawargs |" + rawargs + "|")
		}
		rargs := strings.Split(rawargs, ",")
		for _, v := range rargs {
			if isValidString(v, parser.lexer, tkns[0]) {
				st := utils.GoStrToGagoStr(v)
				args = append(args, ast.Literal{AstType: ast.AST_TYPE_LITERAL, Value: st})
			} else if utils.IsValidInt(v, parser.lexer, tkns[0]) {
				it := utils.GoStrToGagoInt(v)
				args = append(args, ast.Literal{AstType: ast.AST_TYPE_LITERAL, Value: it})
			} else {
				args = append(args, ast.VariableAccess{AstType: ast.AST_TYPE_VARIABLE_ACCESS, Vname: v})
			}
		}
	}

	parser.Ast = append(parser.Ast, ast.FuncCall{AstType: ast.AST_TYPE_FUNC_CALL, Funcname: funcname, Args: args})
}

// TODO: move these functions to a new utils directory
func isValidString(str string, lexer *lexer.Lexer, tkn *lexer.Token) bool {
	qt := 0
	isSqB := string(str[0]) == "'"
	isDqB := string(str[0]) == "\""
	if isSqB {
		qt = 0
	} else if isDqB {
		qt = 1
	} else {
		return false
	}
	isSqE := string(str[len(str)-1]) == "'"
	isDqE := string(str[len(str)-1]) == "\""
	if !(isSqE && qt == 0) && !(isDqE && qt == 1) {
		lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(tkn, lexer.GetFilename()), true).Run()
		return false
	}
	return true
}
