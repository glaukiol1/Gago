package string

import (
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

var FConcat = lang.NewMethod("concat", concatString, "concat joins all the arguments together")
var FContains = lang.NewMethod("contains", containsString, "contains searches for the substring in the specified string. Returns a boolean")
