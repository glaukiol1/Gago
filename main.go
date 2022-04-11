package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	lexer "github.com/glaukiol1/gago/lexer"
	"github.com/glaukiol1/gago/parser"
	"github.com/glaukiol1/gago/vm"
)

// https://jadmogaizel.medium.com/the-different-parts-of-writing-a-programming-language-b634711a6af5

func main() {

	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runFile := runCmd.String("file", "", "file")
	runV := runCmd.Bool("v", false, "file")

	if len(os.Args) < 2 {
		fmt.Println("expected 'run' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "run":
		runCmd.Parse(os.Args[2:])
		run(*runFile, *runV)
	default:
		fmt.Println("Usage: `gago run --file <filename> --v <true|false>")
		os.Exit(1)
	}
}

func run(filename string, v bool) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("fatal error: getting current working directory failed...")
		os.Exit(1)
	}
	file, err := os.Open(path.Join(cwd, filename))
	if err != nil {
		fmt.Println("error while opening file...", err.Error())
		os.Exit(1)
	}
	var fl []byte
	n, err := file.Read(fl)
	if err != nil {
		fmt.Println("error while reading file...", err.Error())
		os.Exit(1)
	}
	if v {
		fmt.Println("read", n, "bytes")
	}
	runfile(string(fl), v)
}

func runfile(filecontents string, v bool) {
	lex := lexer.TestLex(v)
	parse := parser.NewParser(lex)
	parse.Parse()
	vm := vm.NewVM(parse)
	vm.Run()
}
