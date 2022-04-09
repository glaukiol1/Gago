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
	lang.Errorf("SyntaxError", "Unexpected token", lang.BuildStack(tt.token, tt.lexer.GetFilename()), fatal).Run()
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

// isnotspecial checks if a token does not
// represent a special character
// throws a `SyntaxError: Unexpected token` if it is a special character
func (tt *tokentester) IsNotSpecial(fatal bool) bool {
	if !tt.token.IsNewline() && !tt.token.IsWhitespace() && (tt.token.IsCharacter() || tt.token.IsNumber()) {
		return true
	}
	lang.Errorf("SyntaxError", "Unexpected token", lang.BuildStack(tt.token, tt.lexer.GetFilename()), fatal).Run()
	return false
}

// valueis checks if the value of the token
// is equal to `checkv`.
// Throws `SyntaxError: Unexpected indentifier` if it is
// not equal to `checkv`
func (tt *tokentester) ValueIs(checkv interface{}, fatal bool) bool {
	if tt.token.GetValue() != checkv {
		lang.Errorf("SyntaxError", "Unexpected indentifier, expected `"+checkv.(string)+"`", lang.BuildStack(tt.token, tt.lexer.GetFilename()), fatal).Run()
		return false
	}
	return true
}

// nvalueis does the same function as valueis,
// just that it does not throw an error

func (tt *tokentester) NValueIs(checkv interface{}) bool {
	if tt.token.GetValue() != checkv {
		return false
	}
	return true
}
