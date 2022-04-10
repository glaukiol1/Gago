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
	mem.Init(&lang.Options{Stdout: os.Stdout})
	return &VM{v: parser.GetV(), mem: mem, ast: parser.Ast, stdout: os.Stdout}
}

func (vm *VM) Run() {
	for _, v := range vm.ast {
		if ast, ok := v.(ast.VariableDeclaration); ok {
			if vm.v {
				fmt.Println("Running Variable Declaration AST... Vtype: " + fmt.Sprint(ast.Vtype) + " Vname: " + fmt.Sprint(ast.Vname) + " Vvalue: " + ast.Vvalue.Val().(string))
			}
			vm.mem.VarCreate(ast.Vname, ast.Vvalue.(*lang.TypeString)) // FIXME: support more than just TypeString
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
				vm.println("Value: " + val.(string))
			}
		}
	}
}
