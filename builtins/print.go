package builtins

import (
	"fmt"

	"github.com/glaukiol1/gago/lang"
)

// the print() function

func print(args []lang.Type, opt *lang.Options) lang.Type {
	var outtxt string
	for _, t := range args {
		outtxt += fmt.Sprint(t.Val())
	}
	outtxt += "\n"
	opt.Stdout.Write([]byte(outtxt))
	return lang.Null
}

var mprint = lang.NewMethod("print", print, "prints the specified values seperated by a space")
