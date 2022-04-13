package test

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
)

// test module

func Init() *module.Module {
	methods := make(map[string]*lang.Method)
	globals := make(map[string]lang.Type)

	globals["teststring"] = lang.String("hello world!") // test exporting
	return module.NewModule("test", methods, globals)
}
