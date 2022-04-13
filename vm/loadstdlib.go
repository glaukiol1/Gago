package vm

import (
	"github.com/glaukiol1/gago/module"
	"github.com/glaukiol1/gago/stdlib"
)

// load stdlib modules

func LoadStdlib(mem *Memory) {
	for _, v := range stdlib.Modules() {
		loadModuleGlobals(v, mem)
	}
}

func loadModuleGlobals(m *module.Module, mem *Memory) {
	for k, t := range m.GetGlobals() {
		mem.VarCreate(parseModuleName(m, k), t)
	}
	for k, t := range m.GetMethods() {
		mem.LoadMethod(parseModuleName(m, k), t)
	}
}

func parseModuleName(m *module.Module, name string) string {
	return m.GetQualName() + "." + name
}
