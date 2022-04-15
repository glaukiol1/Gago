package string

import "github.com/glaukiol1/gago/lang"

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

var FConcat = lang.NewMethod("concat", concatString, "concat joins all the arguments together")
