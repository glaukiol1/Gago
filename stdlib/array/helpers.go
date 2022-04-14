package array

import (
	"fmt"

	"github.com/glaukiol1/gago/lang"
)

// helper functions for interacting with arrays
// in the Gago programming language

// access will return the item at the specified index or throw a
// IndexError if the item at that index does not exist
func accessArray(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for array.access", true).Run()
	}
	if v, ok := args[0].Val().(*Slice); ok {
		if i, ok := args[1].Val().(int64); ok {
			if i > int64(len(v.Items)-1) {
				lang.Errorf("IndexError", "Index "+fmt.Sprint(i)+" out of bounds.", "", true).Run()
			}
			return v.Items[i]
		} else {
			lang.Errorf("TypeError", "Expected argument of type int (pos 2), but got "+args[1].Name(), "", true).Run()
		}
	} else {
		lang.Errorf("TypeError", "Expected argument of type slice (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return nil
}

// len returns the length of the slice
// on a non-zero based index (starting at 1)
func lenArray(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 1 {
		lang.Errorf("TypeError", "Expected 1 arguments", "\n\t At call for array.len", true).Run()
	}
	if v, ok := args[0].Val().(*Slice); ok {
		return lang.Int(int64(len(v.Items)))
	} else {
		lang.Errorf("TypeError", "Expected argument of type slice (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return nil
}

// pop removes the last entry from a slice
// directly modifies the slice
func popArray(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 1 {
		lang.Errorf("TypeError", "Expected 1 argument", "\n\t At call for array.pop", true).Run()
	}
	if v, ok := args[0].Val().(*Slice); ok {
		v.Items = v.Items[:len(v.Items)-1]
		return lang.Null
	} else {
		lang.Errorf("TypeError", "Expected argument of type slice (pos 1), but got "+args[0].Name(), "", true).Run()
	}
	return nil
}

var FAccess = lang.NewMethod("access", accessArray, "access will return the item at the specified index or throw a IndexError if the item at that index does not exist")
var FLen = lang.NewMethod("len", lenArray, "len returns the length of the slice on a non-zero based index (starting at 1)")
var FPop = lang.NewMethod("pop", popArray, "pop removes the last entry from a slice, directly modifies the slice")
