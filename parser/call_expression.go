package parser

import (
	"fmt"
	"strings"

	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
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
			if isValidString(v) {
				st := goStrToGagoStr(v)
				args = append(args, ast.Literal{AstType: ast.AST_TYPE_LITERAL, Value: st})
			} else {
				args = append(args, ast.VariableAccess{AstType: ast.AST_TYPE_VARIABLE_ACCESS, Vname: v})
			}
		}
	}

	parser.Ast = append(parser.Ast, ast.FuncCall{AstType: ast.AST_TYPE_FUNC_CALL, Funcname: funcname, Args: args})
}

// TODO: move these two function to a new utils directory
func isValidString(str string) bool {
	return (string(str[0]) == "\"" && string(str[len(str)-1]) == "\"") || (string(str[0]) == "'" && string(str[len(str)-1]) == "'")
}

func goStrToGagoStr(str string) *lang.TypeString {
	s := str[1 : len(str)-2]
	return lang.String(s)
}
