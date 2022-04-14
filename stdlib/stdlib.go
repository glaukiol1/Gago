package stdlib

// stdlib package for gago

// this will export all the modules included in the stdlib
// it will help with the creation of the VM since the stdlib
// modules can be imported executed easily

import (
	"github.com/glaukiol1/gago/module"
	array "github.com/glaukiol1/gago/stdlib/array"
	mod "github.com/glaukiol1/gago/stdlib/test"
)

func Modules() []*module.Module {
	var r []*module.Module
	r = append(r, mod.Init())
	r = append(r, array.Init())
	return r
}
