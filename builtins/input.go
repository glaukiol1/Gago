package builtins

// the input function will return a string
// inputs from stdin

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/glaukiol1/gago/lang"
)

// the input() function

func input(args []lang.Type, opt *lang.Options) lang.Type {
	arg := ""
	for _, t := range args {
		arg += fmt.Sprint(t.Val())
	}
	fmt.Print(arg)
	reader := bufio.NewReader(opt.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("internal error: failed to read from stdin")
	}
	str = strings.TrimSuffix(str, "\n")
	return lang.String(str)
}

var minput = lang.NewMethod("input", input, "gets input from stdin until newline")
