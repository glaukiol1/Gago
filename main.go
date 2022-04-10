package main

import (
	lexer "github.com/glaukiol1/gago/lexer"
	"github.com/glaukiol1/gago/parser"
	"github.com/glaukiol1/gago/vm"
)

// https://jadmogaizel.medium.com/the-different-parts-of-writing-a-programming-language-b634711a6af5
func main() {
	lex := lexer.TestLex(false) // change this to false if you dont want to see all the useless output
	parse := parser.NewParser(lex)
	parse.Parse()
	vm := vm.NewVM(parse)
	vm.Run()
}
