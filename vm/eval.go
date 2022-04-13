package vm

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
)

// evaluate expression

// evaluate a variable declaration

// evaluate variable declaration or redeclaration
func eval(x interface{}, vm *VM) lang.Type {
	if _ast, ok := x.(ast.VariableDeclaration); ok {
		return evalexpr(_ast.Vvalue, vm)
	}
	if _ast, ok := x.(ast.VariableReDeclaration); ok {
		return evalexpr(_ast.Vvalue, vm)
	}
	return lang.Null
}

// evaluate a experssion
func evalexpr(arg interface{}, vm *VM) lang.Type {
	if q, ok := arg.(*lang.TypeString); ok {
		return q
	}
	if q, ok := arg.(*lang.TypeInt); ok {
		return q
	}
	if q, ok := arg.(*lang.TypeFloat); ok {
		return q
	}
	if q, ok := arg.(*lang.TypeBool); ok {
		return q
	}
	if q, ok := arg.(ast.Literal); ok {
		return q.Value
	}
	if q, ok := arg.(ast.VariableAccess); ok {
		l, err := vm.mem.AccessVar(q.Vname)
		if err != nil {
			err.(*lang.BaseError).Run()
		}
		return l
	}
	if q, ok := arg.(ast.FuncCall); ok {
		return evalfunc(q, vm) // we may have nested function arguments
	}
	if q, ok := arg.(*ast.MathExpr); ok {
		return evalMathExpr(q.Expression, vm)
	}
	return lang.Null
}

func evalfunc(fc ast.FuncCall, vm *VM) lang.Type {
	mthd, err := vm.mem.AccessMethod(fc.Funcname)
	if err != nil {
		err.(*lang.BaseError).Run()
	}
	var args []lang.Type

	for _, arg := range fc.Args {
		if vm.v {
			fmt.Println("typeof of func arg", reflect.TypeOf(arg).String())
		}
		args = append(args, evalexpr(arg, vm))
	}
	return mthd.RunMethod(args, vm.mem.opts)
}

// eval math equation
func evalMathExpr(s string, vm *VM) lang.Type {
	expr, err := govaluate.NewEvaluableExpression(s)
	if err != nil {
		lang.Errorf("RuntimeError", s+" is not a math expression", "", true).Run()
	}
	args := make(map[string]interface{})
	for k, v := range vm.mem.variables {
		if q, ok := v.Val().(int64); ok {
			if vm.v {
				fmt.Println("parsing variable into math expression...")
			}
			args[k] = q
		}
		if q, ok := v.Val().(float64); ok {
			if vm.v {
				fmt.Println("parsing variable into math expression...")
			}
			args[k] = q
		}
	}
	result, err := expr.Evaluate(args)
	if err != nil {
		fmt.Println(args, s)
		fmt.Println("error while parsing math expression..", err.Error())
		return lang.Null
	}
	if vm.v {
		fmt.Println("result: ", result, "typeof result: ", reflect.TypeOf(result).String())
	}
	r := result.(float64)
	if decimalPortion(r) == 0 {
		return lang.Int(int64(r))
	} else {
		return lang.Float(r)
	}
}

// find out how many decimals a float64 has
func decimalPortion(n float64) int {
	decimalPlaces := fmt.Sprintf("%f", n-math.Floor(n))          // produces 0.xxxx0000
	decimalPlaces = strings.Replace(decimalPlaces, "0.", "", -1) // remove 0.
	decimalPlaces = strings.TrimRight(decimalPlaces, "0")        // remove trailing 0s
	return len(decimalPlaces)
}
