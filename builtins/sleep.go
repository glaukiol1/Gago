package builtins

import (
	"time"

	"github.com/glaukiol1/gago/lang"
)

// this is a non-production ready function

func sleep(args []lang.Type, opt *lang.Options) lang.Type {
	if v, ok := args[0].(*lang.TypeInt); ok {
		time.Sleep(time.Millisecond * time.Duration(v.Val().(int64)))
		return lang.Null
	} else {
		lang.Errorf("TypeError", "expected argument (pos 1) of type int, but got "+v.Name(), "\n\t At call for function sleep()", true).Run()
		return nil
	}
}

var msleep = lang.NewMethod("sleep", sleep, "sleeps for the specified ms")
