package builtins

import (
	"os"
	"reflect"

	"github.com/glaukiol1/gago/lang"
)

// the exit() function

func exit(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) == 0 {
		os.Exit(0)
		return nil
	}
	if v, ok := args[0].Val().(int64); ok {
		os.Exit(int(v))
		return nil
	} else {
		lang.Errorf("TypeError", "Expected argument of type int (pos 1), but instead got "+reflect.TypeOf(args[0].Val()).String(), "", true).Run()
		return nil
	}
}

var mexit = lang.NewMethod("exit", exit, "exits from the process")
