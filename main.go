package main

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gago/parser"
)

func main() {
	fmt.Println("All tests passed: " + strconv.FormatBool(parser.TestLex()))
}
