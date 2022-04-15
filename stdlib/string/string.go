package string

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// the string module is a standard library for
// working with strings in Gago.

func Init() *module.Module {
	methods := make(map[string]*lang.Method)
	globals := make(map[string]lang.Type)

	methods["concat"] = FConcat
	methods["contains"] = FContains
	methods["containsAny"] = FContainsAny
	methods["trimSpace"] = FTrimSpace
	methods["index"] = FIndex
	methods["len"] = FLen

	return module.NewModule("string", methods, globals)
}
