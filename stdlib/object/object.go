package object

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// the object module for Gago
// interaction with Objects in Gago

func Init() *module.Module {
	methods := make(map[string]*lang.Method)
	globals := make(map[string]lang.Type)

	methods["create"] = FCreate
	methods["set"] = FSet
	methods["get"] = FGet
	methods["keys"] = FKeys
	return module.NewModule("object", methods, globals)
}
