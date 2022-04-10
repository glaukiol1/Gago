package parser

import (
	"strconv"

	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// utils specific to the parser

func tokensAreString(cursor *multipleCursor, lexer *lexer.Lexer) bool {
	qt := 0
	for i, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if isSq {
				qt = 0
			} else if isDq {
				qt = 1
			} else {
				return false
			}
		} else if len(cursor.CurrentTokens)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return false
			}
		}
	}
	return true
}

func tokensToGagoString(cursor *multipleCursor, lexer *lexer.Lexer) *lang.TypeString {
	qt := 0
	tmpvalue := ""
	for i, t := range cursor.CurrentTokens {
		tkntest := NewTokenTest(t, lexer)
		if i == 0 {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if isSq {
				qt = 0
			} else if isDq {
				qt = 1
			} else {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return nil
			}
		} else if len(cursor.CurrentTokens)-1 == i {
			isSq := tkntest.NValueIs("'")
			isDq := tkntest.NValueIs("\"")
			if !(isSq && qt == 0) && !(isDq && qt == 1) {
				lang.Errorf("SyntaxError", "Unterminated string literal", lang.BuildStack(cursor.CurrentTokens[0], lexer.GetFilename()), true).Run()
				return nil
			}
		} else {
			tmpvalue += t.GetValue().(string)
		}
	}
	return lang.String(tmpvalue)
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
