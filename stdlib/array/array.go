package array

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// the standard library for interacting with arrays

func Init() *module.Module {
	methods := make(map[string]*lang.Method)
	globals := make(map[string]lang.Type)
	methods["create"] = FCreateArray
	methods["access"] = FAccess
	methods["len"] = FLen
	methods["pop"] = FPop
	methods["shift"] = FShift
	methods["subslice"] = FSubslice
	return module.NewModule("array", methods, globals)
}
