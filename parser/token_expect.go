package parser

import (
	"github.com/glaukiol1/gago/lang"
	"github.com/glaukiol1/gago/lexer"
)

// expect token to be...
// functions to test for conditions with lexer.Tokens

type tokentester struct {
	token *lexer.Token
	lexer *lexer.Lexer
}

func NewTokenTest(token *lexer.Token, lexer *lexer.Lexer) *tokentester {
	return &tokentester{token: token, lexer: lexer}
}

// codeis checks if a tokens code if equal to `match`
// if `fatal` is true, CodeIs throws a fatal SyntaxError
func (tt *tokentester) CodeIs(match int, fatal bool) bool {
	if tt.token.GetCode() == match {
		return true
	}
	lang.Errorf("SyntaxError", "Unexpected token", lang.BuildStack(tt.token, tt.lexer.GetFilename()), fatal)
	return false
}

// ischar checks if a token is a character
// throws the specified `err` if its not a character
func (tt *tokentester) IsChar(err *lang.BaseError) bool {
	if tt.token.IsCharacter() {
		return true
	}
	err.Run()
	return false
}
