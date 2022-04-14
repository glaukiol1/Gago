package object

import "github.com/glaukiol1/gago/lang"

// create a blank object

// creates a blank object
func createObject(args []lang.Type, opt *lang.Options) lang.Type {
	return lang.LoadCustomType("object", &Object{Value: make(map[string]lang.Type)})
}

var FCreate = lang.NewMethod("create", createObject, "creates a blank object")
