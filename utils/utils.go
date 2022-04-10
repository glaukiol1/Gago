package utils

import (
	"strconv"
	"unicode"

	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// this package contains some helper functions
// to make switching between go and gago datatypes
// and other types faster

// note that these are the global ones,
// the other packages might have their own
// utils.go file for their specific needs

func GoStrToGagoStr(str string) *lang.TypeString {
	s := str[1 : len(str)-1]
	return lang.String(s)
}

func IsValidInt(str string, lexer *lexer.Lexer, tkn *lexer.Token) bool {
	for i, q := range str {
		if i == 0 {
			if !(string(q) == "-" || unicode.IsDigit(q)) {
				return false
			}
		} else {
			if !unicode.IsDigit(q) {
				return false
			}
		}
	}
	return true
}

func GoStrToGagoInt(str string) *lang.TypeInt {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil {
		return lang.Int(i)
	}
	return nil
}
