package array

import "github.com/glaukiol1/gago/lang"

// helper functions for interacting with arrays
// in the Gago programming language

// access will return the item at the specified index or throw a
// IndexError if the item at that index does not exist
func accessArray(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for array.access", true).Run()
	}
	if v, ok := args[0].Val().(Slice); ok {
		if i, ok := args[1].Val().(int64); ok {
			return v.Items[i]
		} else {
			lang.Errorf("TypeError", "Expected argument of type int (pos 2), but got "+args[1].Name(), "", true).Run()
		}
	} else {
		lang.Errorf("TypeError", "Expected argument of type slice (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return nil
}

var FAccess = lang.NewMethod("access", accessArray, "access will return the item at the specified index or throw a IndexError if the item at that index does not exist")
