package module

import (
	"os"

	"github.com/glaukiol1/gago/lang"
)

// this file contains the module struct

type Module struct {
	qualanme string                  // qualname
	stdout   *os.File                // stdout
	methods  map[string]*lang.Method // built in methods
	globals  map[string]lang.Type    // built in variables
}

func NewModule(qualname string, methods map[string]*lang.Method, globals map[string]lang.Type) *Module {
	m := &Module{qualanme: qualname, stdout: os.Stdout, methods: methods, globals: globals}
	m.globals = make(map[string]lang.Type) // init globals map
	m.globals = globals
	m.methods = make(map[string]*lang.Method)
	m.methods = methods
	return m
}

func (module *Module) GetMethods() map[string]*lang.Method {
	return module.methods
}

func (module *Module) GetGlobals() map[string]lang.Type {
	return module.globals
}
