package lexer

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	filecontents []byte   // the whole file
	filename     string   // name of the file being read
	eof          bool     // flag to show EOF was read
	bracket      int      // number of brackets
	parenthesis  int      // number of parenthesis
	curly_braces int      // number of curly braces
	tokens       []*Token // buffered tokens to output
	error        bool     // set if an error has ocurred
	errorString  string   // the string of the error
}

func NewLex(filecontent, filename string) *Lexer {
	return &Lexer{[]byte(filecontent), filename, false, 0, 0, 0, []*Token{}, false, ""}
}

func (lexer *Lexer) GetTokens() []*Token {
	return lexer.tokens
}

func (lexer *Lexer) GetFilename() string {
	return lexer.filename
}

func (lexer *Lexer) Lex(v bool) error {
	str_file := string(lexer.filecontents)
	for pos, ln := range str_file {
		if pos == len(str_file)-1 {
			lexer.tokens = append(lexer.tokens, NewToken(string(ln), pos), NewToken("EOF", pos))
			break
		}
		token := NewToken(string(ln), pos)
		if token.code == 13 && lexer.tokens[len(lexer.tokens)-1].code == 69 && lexer.tokens[len(lexer.tokens)-2].code != 69 {
			lexer.tokens[len(lexer.tokens)-1] = NewToken("\n", pos)
		} else {
			lexer.tokens = append(lexer.tokens, token)
		}
	}
	if v {
		fmt.Println(lexer.bracket, lexer.curly_braces, lexer.parenthesis)
		for _, token := range lexer.tokens {
			fmt.Println("Token at line " + fmt.Sprint(token.line) + ", at pos " + fmt.Sprint(token.pos) + " value: " + fmt.Sprint(token.value) + " code: " + fmt.Sprint(token.code) + " Is character: " + strconv.FormatBool(token.IsCharacter()) + " Is whitespace: " + strconv.FormatBool(token.IsWhitespace()) + " Is Number " + strconv.FormatBool(token.IsNumber()))
		}
	}
	return nil
}
