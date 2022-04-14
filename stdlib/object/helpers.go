package object

import (
	"github.com/glaukiol1/gago/lang"
)

// functions for interacting with objects
// in Gago

// set a value in the given object.
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

// get a value in the given object.
func getObject(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for object.get", true).Run()
	}
	if v, ok := args[0].Val().(*Object); ok {
		if i1, ok := args[1].Val().(string); ok {
			return v.Value[i1]
		} else {
			lang.Errorf("TypeError", "Expected argument of type string (pos 2), but got "+args[1].Name(), "", true).Run()
		}
	} else {
		lang.Errorf("TypeError", "Expected argument of type object (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return lang.Null
}

var FSet = lang.NewMethod("set", setObject, "set a value in the given object.")
var FGet = lang.NewMethod("get", getObject, "get a value in the given object.")
