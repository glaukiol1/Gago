package lang

// method is a special Type, which will not inherit
// from the basetype.

// a gago function that is written in Go will look something like this
/*
func testgagofunction(args []lang.Type) lang.Type {
	return args[0]
}
*/
// this function will return the first argument
// that it was given.

type Method struct {
	methodname string                 // the method name inside the gago script
	gomethod   func(args []Type) Type // the method which will be called
	doc        string                 // the doc for this function
}

func NewMethod(methodname string, gomethod func(args []Type) Type, doc string) *Method {
	return &Method{methodname: methodname, gomethod: gomethod, doc: doc}
}

func (method *Method) RunMethod(args []Type) Type {
	return method.gomethod(args)
}
