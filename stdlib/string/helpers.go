package string

import (
	"fmt"
	"strings"

	"github.com/glaukiol1/gago/lang"
)

// this file contains helper functions
// for interacting with strings

// concat joins all the arguments together
func concatString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) < 2 {
		lang.Errorf("TypeError", "Expected 2 or more arguments", "\n\t At call for string.concat", true).Run()
	}
	result := ""
	for _, x := range args {
		if x.Name() == "string" {
			result += x.Val().(string)
		} else {
			lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.concat", true).Run()
		}
	}
	return lang.String(result)
}

// contains searches for the substring in the specified
// string. Returns a boolean
func containsString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for string.contains", true).Run()
	}
	str := ""
	searchstr := ""
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.contains", true).Run()
	}

	x = args[1]
	if x.Name() == "string" {
		searchstr += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.contains", true).Run()
	}
	res := strings.Contains(str, searchstr)
	if res == true {
		return lang.True
	}
	return lang.False
}

// containsAny checks if any of the chars are in
// the set string. Returns a boolean
func containsAnyString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for string.contains", true).Run()
	}
	str := ""
	searchstr := ""
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.containsAny", true).Run()
	}

	x = args[1]
	if x.Name() == "string" {
		searchstr += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.containsAny", true).Run()
	}
	res := strings.ContainsAny(str, searchstr)
	if res == true {
		return lang.True
	}
	return lang.False
}

// trimSpace removes all trailing and leading whitespaces
// and returns the new string
func trimSpaceString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 1 {
		lang.Errorf("TypeError", "Expected 1 argument", "\n\t At call for string.trimSpace", true).Run()
	}
	str := ""
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.trimSpace", true).Run()
	}
	return lang.String(strings.TrimSpace(str))
}

// Index returns the index of the first instance of substr
// in string, or -1 if substr is not present in string.
func indexString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for string.index", true).Run()
	}
	str := ""
	searchstr := ""
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.index", true).Run()
	}

	x = args[1]
	if x.Name() == "string" {
		searchstr += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.index", true).Run()
	}
	res := strings.Index(str, searchstr)
	return lang.Int(int64(res))
}

// len returns the length of the string
func lenString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 1 {
		lang.Errorf("TypeError", "Expected 1 argument", "\n\t At call for string.len", true).Run()
	}
	str := ""
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.len", true).Run()
	}
	return lang.Int(int64(len(str)))
}

// charAt returns the character
// at the given index, or a IndexError
func charAtString(args []lang.Type, opt *lang.Options) lang.Type {
	if len(args) != 2 {
		lang.Errorf("TypeError", "Expected 2 arguments", "\n\t At call for string.charAt", true).Run()
	}
	str := ""
	var idx int64
	x := args[0]
	if x.Name() == "string" {
		str += x.Val().(string)
	} else {
		lang.Errorf("TypeError", "Expected argument of type string", "\n\t At call for string.charAt", true).Run()
	}

	x = args[1]
	if x.Name() == "int" {
		idx = x.Val().(int64)
	} else {
		lang.Errorf("TypeError", "Expected argument of type int64", "\n\t At call for string.charAt", true).Run()
	}
	if idx > int64(len(str))-1 {
		lang.Errorf("IndexError", "Index "+fmt.Sprint(idx)+" out of bounds", "\n\t At call for string.charAt", true).Run()
	}
	return lang.String(string(str[idx]))
}

var FConcat = lang.NewMethod("concat", concatString, "concat joins all the arguments together")
var FContains = lang.NewMethod("contains", containsString, "contains searches for the substring in the specified string. Returns a boolean")
var FContainsAny = lang.NewMethod("containsAny", containsAnyString, "containsAny checks if any of the chars are in the set string. Returns a boolean")
var FTrimSpace = lang.NewMethod("trimSpace", trimSpaceString, "trimSpace removes all trailing and leading whitespaces and returns the new string")
var FIndex = lang.NewMethod("index", indexString, "index returns the index of the first instance of substr in string, or -1 if substr is not present in string.")
var FLen = lang.NewMethod("len", lenString, "len returns the length of the string")
var FCharAt = lang.NewMethod("charAt", charAtString, "charAt returns the character at the given index, or a IndexError")
