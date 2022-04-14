package run

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/glaukiol1/gago/lexer"
	"github.com/glaukiol1/gago/parser"
	"github.com/glaukiol1/gago/vm"
)

// run file or string with Gago

func RunFile(filename string, v bool) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("fatal error: getting current working directory failed...")
		os.Exit(1)
	}
	content, err := os.ReadFile(path.Join(cwd, filename))
	if err != nil {
		log.Fatal(err)
	}
	RunData(string(content), filename, v)
}

func RunData(filecontents string, flname string, v bool) {
	lex := lexer.NewLex(filecontents, flname, v)
	lex.Lex()
	parse := parser.NewParser(lex)
	parse.Parse()
	vm := vm.NewVM(parse)
	vm.Run()
}
