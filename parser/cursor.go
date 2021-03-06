package parser

import (
	"github.com/glaukiol1/gago/lexer"
)

type multipleCursor struct {
	CurrentTokens []*lexer.Token   // currentToken
	currentIndex  int              // index of current token
	tokens        [][]*lexer.Token // all tokens
	linepos       int              // line position
}

// init a new cursor
func newMultipleCursor(tokens [][]*lexer.Token, linepos int) *multipleCursor {
	// this loop sets the line of the token
	for _, x := range tokens {
		for _, z := range x {
			z.SetLine(linepos + 1) // +1 because the index starts at 0, white the file starts at 1
		}
	}
	return &multipleCursor{tokens: tokens, CurrentTokens: tokens[0], currentIndex: 0, linepos: linepos}
}

// switch to the next set of tokens
// returns is there is a next token or not
func (cursor *multipleCursor) Next() bool {
	cursor.currentIndex += 1
	if len(cursor.tokens) == cursor.currentIndex {
		return false
	}
	cursor.CurrentTokens = cursor.tokens[cursor.currentIndex]
	return true
}

// set to a index inside the cursor slice
func (cursor *multipleCursor) SetIndex(index int) {
	cursor.currentIndex = index

	cursor.CurrentTokens = cursor.tokens[cursor.currentIndex]
}

// join all from will join all from the specified index
func (cursor *multipleCursor) JoinAllFrom(indx int, sep string) []*lexer.Token {
	var rtrn []*lexer.Token
	flag := false
	sepToken := lexer.NewToken(sep, cursor.CurrentTokens[0].GetLine())
	for i, v := range cursor.tokens {
		if i == indx || flag {
			flag = true
			rtrn = append(rtrn, sepToken)
			rtrn = append(rtrn, v...)
		}
	}
	return rtrn
}

func (cursor *multipleCursor) SubCursor(indx int) *multipleCursor {
	return newMultipleCursor(cursor.tokens[indx:], cursor.linepos)
}

func (cursor *multipleCursor) Before() {
	cursor.currentIndex -= 1
	cursor.CurrentTokens = cursor.tokens[cursor.currentIndex]
}
