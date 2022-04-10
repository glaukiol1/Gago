package vm

import (
	"fmt"

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
	v   bool          // verbose
	mem *Memory       // the memory of the VM | defined in memory.go
	ast []interface{} // interface{} will be one of the structs in `ast` package
}

func NewVM(parser *parser.Parser) *VM {
	return &VM{v: parser.GetV(), mem: NewMemory(parser.GetV()), ast: parser.Ast}
}

func (vm *VM) Run() {
	for _, v := range vm.ast {
		if ast, ok := v.(ast.VariableDeclaration); ok {
			if vm.v {
				fmt.Println("Running AST... Vtype: " + fmt.Sprint(ast.Vtype) + " Vname: " + fmt.Sprint(ast.Vname) + " Vvalue: " + ast.Vvalue.Val().(string))
			}
			vm.mem.VarCreate(ast.Vname, ast.Vvalue.(*lang.TypeString))
		}
	}
}
