package lang

import "os"

// method is a special Type, which will not inherit
// from the basetype.

// a gago function that is written in Go will look something like this
/*
func testgagofunction(args []lang.Type, options *lang.Options) lang.Type {
	return args[0]
}
*/
// this function will return the first argument
// that it was given.

type Method struct {
	qualname string                                // the method name inside the gago script
	gomethod func(args []Type, opts *Options) Type // the method which will be called
	doc      string                                // the doc for this function
}

type Options struct {
	Stdout *os.File
}

func NewMethod(qualname string, gomethod func(args []Type, opts *Options) Type, doc string) *Method {
	return &Method{qualname: qualname, gomethod: gomethod, doc: doc}
}

func (method *Method) RunMethod(args []Type, opts *Options) Type {
	return method.gomethod(args, opts)
}

func (method *Method) GetDoc() string {
	return method.doc
}

func (method *Method) GetName() string {
	return method.qualname
}
