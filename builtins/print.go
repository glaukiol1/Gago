package builtins

import (
	"fmt"

	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/stdlib/array"
)

// the print() function

func print(args []lang.Type, opt *lang.Options) lang.Type {
	var outtxt string
	for _, t := range args {
		if v, ok := t.Val().(*array.Slice); ok {
			outtxt += "["
			for i, t2 := range v.Items {
				outtxt += fmt.Sprint(t2.Val())
				if i != len(v.Items)-1 {
					outtxt += ", "
				}
			}
			outtxt += "]"
		} else {
			outtxt += fmt.Sprint(t.Val())
		}
		outtxt += " "
	}
	outtxt += "\n"
	opt.Stdout.Write([]byte(outtxt))
	return lang.Null
}

var mprint = lang.NewMethod("print", print, "prints the specified values seperated by a space")
