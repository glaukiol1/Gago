package parser

import (
	"strconv"
	"strings"

	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// utils specific to the parser

func tokensString(cursor *multipleCursor, lexer *lexer.Lexer) (lang.Type, bool) {
	qt := 0
	tkns := cursor.JoinAllFrom(3, " ")
	tmpvalue := ""
	i := 0
	// this loop removes all leading whitespaces
	if len(tkns) == 0 {
		return nil, false
	}
	for {
		if tkns[i].IsWhitespace() {
			i += 1
		} else {
			break
		}
	}
	//
	tkns = tkns[i:] // new array where the whitespace leading tokens are removed
	for i, t := range tkns {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if isSq {
				qt = 0
			} else if isDq {
				qt = 1
			} else {
				return nil, false
			}
		} else if len(tkns)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(tkns[i], lexer.GetFilename()), true).Run()
				return nil, false
			}
		} else {
			tmpvalue += t.GetValue().(string)
		}
	}
	return lang.String(tmpvalue), true
}

func tokensAreInt(cursor *multipleCursor, lexer *lexer.Lexer) bool {
	for i, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			if !(tkntest.NValueIs("-") || tkntest.IsNum()) {
				return false
			}
		} else {
			if !tkntest.IsNum() {
				return false
			}
		}
	}
	return true
}

func tokensToGagoInt(cursor *multipleCursor, lexer *lexer.Lexer) *lang.TypeInt {
	var str string
	for _, v := range cursor.CurrentTokens {
		q, ok := v.GetValue().(string)
		if !ok {
			panic("internal error: tokenstogagoint failed")
		}
		str += q
	}
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("internal error: tokenstogagoint failed")
	}
	return lang.Int(n)
}

// checks if cursor.CurrentTokens can represent a float
func tokensAreFloat(cursor *multipleCursor, lexer *lexer.Lexer) bool {
	var str string
	for _, v := range cursor.CurrentTokens {
		q, ok := v.GetValue().(string)
		if !ok {
			panic("internal error: tokenstogagoint failed")
		}
		str += q
	}
	dotsplit := strings.Split(str, ".")
	if len(dotsplit) != 2 {
		return false
	}
	_, err := strconv.ParseInt(dotsplit[0], 10, 64)
	if err != nil {
		return false
	}
	_, err = strconv.ParseInt(dotsplit[1], 10, 64)
	if err != nil {
		return false
	}
	_, err = strconv.ParseFloat(str, 64)
	return err == nil
}

func tokensToFloat(cursor *multipleCursor, lexer *lexer.Lexer) *lang.TypeFloat {
	var str string
	for _, v := range cursor.CurrentTokens {
		q, ok := v.GetValue().(string)
		if !ok {
			panic("internal error: tokenstogagoint failed")
		}
		str += q
	}
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("internal error: failed to parse float in tokensToFloat")
	}
	return lang.Float(n)
}
