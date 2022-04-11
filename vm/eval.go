package vm

import (
	"fmt"
	"reflect"

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
