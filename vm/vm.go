package vm

import (
	"fmt"
	"os"

	"github.com/glaukiol1/gago/ast"
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/parser"
)

// vm | this package contains the virtual machine
// for running gago code

// The VM struct will be the holder of all the
// subcomponents of the VM
type VM struct {
	v      bool          // verbose
	mem    *Memory       // the memory of the VM | defined in memory.go
	ast    []interface{} // interface{} will be one of the structs in `ast` package
	stdout *os.File      // standard out of gago
}

func NewVM(parser *parser.Parser) *VM {
	mem := NewMemory(parser.GetV())
	mem.Init(&lang.Options{Stdout: os.Stdout, Stdin: os.Stdin})
	return &VM{v: parser.GetV(), mem: mem, ast: parser.Ast, stdout: os.Stdout}
}

func (vm *VM) Run() {
	for _, v := range vm.ast {
		if _ast, ok := v.(ast.VariableDeclaration); ok {
			if vm.v {
				fmt.Println("Running Variable Declaration AST... Vtype: "+fmt.Sprint(_ast.Vtype)+" Vname: "+fmt.Sprint(_ast.Vname)+" Vvalue: ", _ast.Vvalue)
			}
			vm.mem.VarCreate(_ast.Vname, eval(_ast, vm))
		}

		// -- in the case of a nested ast.VariableAccess --
		// ast.VariableAccess will just subsitute the AST with the variable value,
		// so if a function is called, example code:
		// ---
		// var testvar = "Hello World"
		// print(testvar)
		// ---
		// the `testvar` will be a ast.VariableAccess, and then the below code
		// will subsitute `testvar` with the value of it, in this case, `Hello World`.
		// it will subsitute it not with a Go datatype, but a *lang.Type (which can be any of the
		// goga datatypes).
		// ast.VariableAccess will be a nested AST.

		// the non-nested ast.Variable Access can not happen, and is only for debug purposes
		if ast, ok := v.(ast.VariableAccess); ok {
			if vm.v {
				fmt.Println("Running Variable Access AST... Vname: " + ast.Vname)
			}
			val, err := vm.mem.AccessVar(ast.Vname)
			if err != nil {
				err.(*lang.BaseError).Run()
				return
			}
			if vm.v {
				vm.println("Value: " + val.Val().(string))
			}
		}

		// if the ast is of type ast.FuncCall, use the memory to access
		// and run the named method
		// note that this is not a nested call.
		if ast_, ok := v.(ast.FuncCall); ok {
			evalfunc(ast_, vm)
		}
	}
}

func (vm *VM) SetAST(ast []interface{}) {
	vm.ast = ast
}
