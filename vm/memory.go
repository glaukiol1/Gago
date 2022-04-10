package vm

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gago/builtins"
	"github.com/glaukiol1/gago/lang"
)

// this file will contain all actions
// related to memory in the gago VM

type Memory struct {
	v         bool                    // verbose
	variables map[string]lang.Type    // this map will hold all variables
	methods   map[string]*lang.Method // this map holds all the global methods
	opts      *lang.Options           // options
	// modules []*module.Module TODO: add module.Module
}

func NewMemory(v bool) *Memory {
	m := &Memory{v: true, methods: nil}
	m.variables = make(map[string]lang.Type)  // initialize variable map
	m.methods = make(map[string]*lang.Method) // initialize method map
	return m
}

func (mem *Memory) LoadBuiltins() {
	bmodule := builtins.FBuilitins(mem.opts.Stdout)
	for k, v := range bmodule.GetGlobals() {
		if mem.v {
			fmt.Println("global builtin: k: |" + k + "| v: |" + v.Val().(string) + "|")
		}
		mem.VarCreate(k, v)
	}
	for k, v := range bmodule.GetMethods() {
		mem.LoadMethod(k, v)
	}
	if mem.v {
		fmt.Println("successfully loaded builtins")
	}
}

func (mem *Memory) Init(opts *lang.Options) {
	mem.opts = opts
	mem.LoadBuiltins()
}

func (mem *Memory) VarCreate(name string, value interface{}) {
	if val, ok := value.(lang.Type); ok {
		mem.variables[name] = val
		if mem.v {
			fmt.Println("Added variable to memory... Name: " + name + " Value: " + val.Val().(string) + " Constant: " + strconv.FormatBool(val.IsConstant()))
		}
	}
}

func (mem *Memory) VarExists(name string) (lang.Type, bool) {
	for k, v := range mem.variables {
		if k == name {
			return v, true
		}
	}
	return nil, false
}

func (mem *Memory) VarUpdate(name string, value interface{}) error {
	if t, ok := mem.VarExists(name); ok {
		if t.IsConstant() {
			return lang.Errorf("TypeError", "Assignment to constant variable.", "At variable "+t.Name(), true)
		}
		t.Reassign(value)
		return nil
	}
	return lang.Errorf("RuntimeError", "Unable to reasssign to variable "+name, "", true)
}

func (mem *Memory) AccessVar(name string) (interface{}, error) {
	if t, ok := mem.VarExists(name); ok {
		return t.Val(), nil
	}
	return nil, lang.Errorf("ReferenceError", name+" is not defined.", "", true)
}

func (mem *Memory) LoadMethod(name string, method *lang.Method) {
	mem.methods[name] = method
}

func (mem *Memory) AccessMethod(qualname string) (*lang.Method, error) {
	if mthd, ok := mem.MethodExists(qualname); ok {
		return mthd, nil
	}
	return nil, lang.Errorf("ReferenceError", qualname+" is not defined.", "", true)
}

func (mem *Memory) MethodExists(qualname string) (*lang.Method, bool) {
	for k, v := range mem.methods {
		if k == qualname {
			return v, true
		}
	}
	return nil, false
}
