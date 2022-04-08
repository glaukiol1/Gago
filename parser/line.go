package parser

import (
	"fmt"
	"os"
)

type Line struct {
	lexer        *Lexer
	line         string    // line text
	linepos      int       // line number
	cursors      []*Cursor // Cursors in this line
	indents      int       // indents in current line
	error        bool      // set if an error has ocurred
	errorString  string    // the string of the error
	bracket      int       // number of brackets
	parenthesis  int       // number of parenthesis
	curly_braces int       // number of curly braces
}

func NewLine(lexer *Lexer, line string, linepos int) *Line {
	return &Line{lexer, line, linepos, []*Cursor{}, 0, false, "", 0, 0, 0}
}

func (line *Line) ParseLine() {
	for pos, char := range line.line {
		cursor, err := NewCursor(line, string(char), pos)
		if err != nil {
			fmt.Println("Error while tokenizing: " + line.errorString)
			os.Exit(1)
		}
		line.cursors = append(line.cursors, cursor)
	}
	line.syncWithLexer()
}

func (line *Line) syncWithLexer() {
	if line.error {
		fmt.Println(line.errorString)
		os.Exit(1)
	}
	line.lexer.tokens = append(line.lexer.tokens, line.toTokens()...)
	line.lexer.error = line.error
	line.lexer.errorString = line.errorString
	line.lexer.bracket += line.bracket
	line.lexer.curly_braces += line.curly_braces
	line.lexer.parenthesis += line.parenthesis
}

func (line *Line) IsEmpty() bool {
	if len(line.cursors) == 0 {
		return true
	}
	return false
}

func (line *Line) toTokens() []*Token {
	var tokens []*Token
	for _, cursor := range line.cursors {
		tokens = append(tokens, cursor.Token)
	}
	return tokens
}
