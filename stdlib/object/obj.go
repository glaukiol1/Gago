package object

import "github.com/glaukiol1/gago/lang"

// obj datatype

type Object struct {
	Value map[string]lang.Type // the value of the object
}
