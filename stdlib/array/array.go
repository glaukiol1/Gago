package array

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// the standard library for interacting with arrays

func Init() *module.Module {
	return module.NewModule("array", make(map[string]*lang.Method), make(map[string]lang.Type))
}
