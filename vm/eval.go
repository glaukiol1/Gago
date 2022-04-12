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

func eval(_ast ast.VariableDeclaration, vm *VM) lang.Type {
	if q, ok := _ast.Vvalue.(*lang.TypeString); ok {
		q.SetConstant(_ast.Vtype == 1)
		return q
	}
	if q, ok := _ast.Vvalue.(*lang.TypeInt); ok {
		q.SetConstant(_ast.Vtype == 1)
		return q
	}
	if q, ok := _ast.Vvalue.(ast.FuncCall); ok {
		v := evalfunc(q, vm)
		v.SetConstant(_ast.Vtype == 1)
		return v
	}
	if q, ok := _ast.Vvalue.(ast.VariableAccess); ok {
		v, err := vm.mem.AccessVar(q.Vname)
		if err != nil {
			err.(*lang.BaseError).Run()
		}
		return v
	}
	if q, ok := _ast.Vvalue.(*ast.MathExpr); ok {
		v := evalMathExpr(q.Expression, vm)
		v.SetConstant(_ast.Vtype == 1)
		return v
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
		if q, ok := arg.(*lang.TypeString); ok {
			args = append(args, q)
		}
		if q, ok := arg.(*lang.TypeInt); ok {
			args = append(args, q)
		}
		if q, ok := arg.(ast.Literal); ok {
			args = append(args, q.Value)
		}
		if q, ok := arg.(ast.VariableAccess); ok {
			l, err := vm.mem.AccessVar(q.Vname)
			if err != nil {
				err.(*lang.BaseError).Run()
			}
			args = append(args, l)
		}
		if q, ok := arg.(ast.FuncCall); ok {
			args = append(args, evalfunc(q, vm)) // we may have nested function arguments
		}
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
	}
	result, err := expr.Evaluate(args)
	if vm.v {
		fmt.Println("result: ", result, "typeof result: ", reflect.TypeOf(result).String())
	}
	if err != nil {
		fmt.Println("error while parsing math expression..", err.Error())
		return lang.Null
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
