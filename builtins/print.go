package builtins

import "github.com/glaukiol1/gago/lang"

// the print() function

func print(args []lang.Type, opt *lang.Options) lang.Type {
	var outtxt string
	for i, t := range args {
		if v, ok := t.(*lang.TypeString); ok {
			if i != 0 {
				outtxt += " "
			}
			outtxt += v.Val().(string)
		}
	}
	outtxt += "\n"
	opt.Stdout.Write([]byte(outtxt))
	return &lang.Null
}

var mprint = lang.NewMethod("print", print, "prints the specified values seperated by a space")
