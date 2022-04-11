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

// TODO
// there is a virtual memory map
// which keeps track of variables and other things

// TODO
// there is a way to load modules,
// and make their globals accessable

// TODO
// runs the AST, and is able to parse
// and run the according functions
// corresponding to the AST

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
		if ast, ok := v.(ast.VariableDeclaration); ok {
			if vm.v {
				fmt.Println("Running Variable Declaration AST... Vtype: "+fmt.Sprint(ast.Vtype)+" Vname: "+fmt.Sprint(ast.Vname)+" Vvalue: ", ast.Vvalue.Val())
			}
			if q, ok := ast.Vvalue.(*lang.TypeString); ok {
				vm.mem.VarCreate(ast.Vname, q)
			}
			if q, ok := ast.Vvalue.(*lang.TypeInt); ok {
				vm.mem.VarCreate(ast.Vname, q)
			}
		}

		// in the future, ast.VariableAccess will just subsitute the AST with the variable value,
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
		if ast_, ok := v.(ast.FuncCall); ok {
			var args []lang.Type
			for _, r := range ast_.Args {

				// check if the Argument is of type VariableAccess
				// meaning that it is a variable name, not a literal
				if _ast, ok := r.(ast.VariableAccess); ok {
					vval, err := vm.mem.AccessVar(_ast.Vname)
					if err != nil {
						err.(*lang.BaseError).Run()
					}
					args = append(args, vval)
				}

				// check if the Argument is of type Literal
				// meaning that it is not a variable name,
				// but a literal
				if _ast, ok := r.(ast.Literal); ok {
					args = append(args, _ast.Value)
				}
			}
			mthd, err := vm.mem.AccessMethod(ast_.Funcname)
			if err != nil {
				err.(*lang.BaseError).Run()
			}
			mthd.RunMethod(args, vm.mem.opts)
		}
	}
}
