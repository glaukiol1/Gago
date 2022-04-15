package stdlib

// stdlib package for gago

// this will export all the modules included in the stdlib
// it will help with the creation of the VM since the stdlib
// modules can be imported executed easily

import (
	"github.com/glaukiol1/gago/module"
	array "github.com/glaukiol1/gago/stdlib/array"
	object "github.com/glaukiol1/gago/stdlib/object"
	string "github.com/glaukiol1/gago/stdlib/string"
	mod "github.com/glaukiol1/gago/stdlib/test"
)

func Modules() []*module.Module {
	var r []*module.Module
	r = append(r, mod.Init())
	r = append(r, array.Init())
	r = append(r, object.Init())
	r = append(r, string.Init())
	return r
}
