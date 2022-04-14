package array

import (
	"github.com/glaukiol1/gago/lang"
)

// create an array function

func createArray(args []lang.Type, opt *lang.Options) lang.Type {
	var arrayValue []interface{}
	for _, arg := range args {
		arrayValue = append(arrayValue, arg.Val())
	}
	return lang.LoadCustomType("slice", arrayValue)
}

var FCreateArray = lang.NewMethod("create", createArray, "Create a new array slice of all the argumentss")
