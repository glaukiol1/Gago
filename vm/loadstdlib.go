package vm

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/module"
	"github.com/glaukiol1/gago/stdlib"
)

// load stdlib modules

func LoadStdlib(mem *Memory, name string) {
	for _, v := range stdlib.Modules() {
		if v.GetQualName() == name {
			loadModule(v, mem)
			return
		}
	}
	lang.Errorf("ImportError", "Couldn't locate module `"+name+"`", "", true).Run()
}

func loadModule(m *module.Module, mem *Memory) {
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
