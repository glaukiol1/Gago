package object

import (
	"github.com/glaukiol1/gago/lang"
)

// functions for interacting with objects
// in Gago

// init a empty object
func setObject(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 3 {
		lang.Errorf("TypeError", "Expected 3 arguments", "\n\t At call for object.set", true).Run()
	}
	if v, ok := args[0].Val().(*Object); ok {
		if i1, ok := args[1].Val().(string); ok {
			v.Value[i1] = args[2]
			return lang.Null
		} else {
			lang.Errorf("TypeError", "Expected argument of type string (pos 2), but got "+args[1].Name(), "", true).Run()
		}
	} else {
		lang.Errorf("TypeError", "Expected argument of type object (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return lang.Null
}

var FSet = lang.NewMethod("set", setObject, "init a empty object")
