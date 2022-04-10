package vm

import (
	"fmt"
	"strconv"

	"github.com/glaukiol1/gago/lang"
)

// this file will contain all actions
// related to memory in the gago VM

type Memory struct {
	v         bool                 // verbose
	variables map[string]lang.Type // this slice will hold all variables
	// methods []*lang.Method TODO: add lang.Method
	// modules []*module.Module TODO: add module.Module
}

func NewMemory(v bool) *Memory {
	m := &Memory{v: true}
	m.variables = make(map[string]lang.Type) // initialize map
	return m
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
