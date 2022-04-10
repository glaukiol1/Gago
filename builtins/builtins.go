package builtins

import (
	"os"

	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// builtin globals and functions

func FBuilitins(stdout *os.File) *module.Module {
	methods := make(map[string]*lang.Method)
	globals := make(map[string]lang.Type)

	methods["print"] = mprint
	globals["teststring"] = lang.String("hello world!") // test exporting builtins
	return module.NewModule("builtins", methods, globals)
}