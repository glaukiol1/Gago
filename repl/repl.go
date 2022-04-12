package repl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/glaukiol1/gago/lexer"
	"github.com/glaukiol1/gago/parser"
	"github.com/glaukiol1/gago/vm"
)

// repl | Gago
// the repl is the interactive terminal where
// a user can enter Gago code and it will be lexed,
// parsed, and ran in the VM.

func Start(version string) {
	fmt.Println("Gago Repl")
	fmt.Printf("[Gago %s]\n", version)
	fmt.Printf("- os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- go version: %s\n", runtime.Version())

	vm := vm.NewVM(&parser.Parser{})
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("internal error: failed to read from stdin")
		}
		xast := lexparse(str)
		vm.SetAST(xast)
		vm.Run()
	}
}

// lexparse returns the AST of the f string
func lexparse(f string) []interface{} {
	lex := lexer.NewLex(f, "<repl>", false)
	lex.Lex()
	parse := parser.NewParser(lex)
	parse.Parse()
	return parse.Ast
}
