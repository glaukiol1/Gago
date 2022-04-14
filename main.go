package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glaukiol1/gago/repl"
	"github.com/glaukiol1/gago/run"
)

// https://jadmogaizel.medium.com/the-different-parts-of-writing-a-programming-language-b634711a6af5

func main() {

	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runFile := runCmd.String("file", "", "file")
	runV := runCmd.Bool("v", false, "file")

	if len(os.Args) == 1 {
		repl.Start("v0.5-alpha")
	}

	if len(os.Args) < 2 {
		fmt.Println("expected 'run' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		run.RunFile(*runFile, *runV)
	default:
		fmt.Println("Usage: `gago run --file <filename> --v <true|false>")
		os.Exit(1)
	}
}
