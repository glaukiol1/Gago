package parser

import (
	"fmt"
	"reflect"

	"github.com/glaukiol1/gago/lexer"
)

type Parser struct {
	lexer         *lexer.Lexer     // lexer
	tokens        []*lexer.Token   // tokens
	newlineTokens [][]*lexer.Token // tokens seperated by whitesapces
	error         bool             // keep track of errors
	errorString   string           // error string
	v             bool             // verbose
}

func NewParser(lexer *lexer.Lexer, v bool) *Parser {
	return &Parser{lexer: lexer, tokens: lexer.GetTokens(), error: false, errorString: "", v: v}
}

func (parser *Parser) sepNewlines() {
	var local []*lexer.Token
	for _, t := range parser.tokens {
		if t.IsNewline() {
			parser.newlineTokens = append(parser.newlineTokens, local)
			local = []*lexer.Token{}
		} else if t.GetCode() == 96 { // EOF
			parser.newlineTokens = append(parser.newlineTokens, local)
		} else {
			local = append(local, t)
		}
	}
}

func (parser *Parser) sepWhiteSpaces(newlinetokens []*lexer.Token) [][]*lexer.Token {
	var array [][]*lexer.Token
	var local []*lexer.Token
	for i, token := range newlinetokens {
		if i == len(newlinetokens)-1 {
			local = append(local, token)
			array = append(array, local)
			break
		}
		if token.GetCode() == 94 {
			if len(local) != 0 {
				array = append(array, local)
				local = []*lexer.Token{}
			}
		} else {
			local = append(local, token)
		}
	}
	return array
}

func (parser *Parser) Parse(v bool) {
	parser.sepNewlines()
	if v {
		for linepos, newline_sep_tokens := range parser.newlineTokens {
			fmt.Println("---" + fmt.Sprint(linepos) + "---")
			rmvd, _ := parser.removeTrailingSpaces(newline_sep_tokens)
			whitespace_sep_tokens := parser.sepWhiteSpaces(rmvd)
			for _, v := range whitespace_sep_tokens {
				parser.parsenewlineTokens(v)
			}
		}
	}
}

// parse sep whitespace tokens inside a newline
// Example: [token("c"),token("o"),token("n"),token("s"),token("t")]
func (parser *Parser) parsenewlineTokens(whitespace_sep_tokens []*lexer.Token) {
	var codes []int
	for _, t := range whitespace_sep_tokens {
		codes = append(codes, t.GetCode())
	}
	parser.checkPattern(codes)
}

// match token codes with known patterns
func (parser *Parser) checkPattern(codes []int) interface{} {
	if reflect.DeepEqual(codes, KEYWORD_CONST_CODE) {
		if parser.v {
			fmt.Println("found const statement")
		}
		return nil
	} else if reflect.DeepEqual(codes, KEYWORD_VAR_CODE) {
		if parser.v {
			fmt.Println("found var statement")
		}
		return nil
	}
	return nil
}

func (parser *Parser) removeTrailingSpaces(whitespace_sep_tokens []*lexer.Token) ([]*lexer.Token, int) {
	cursor := 0
	whitespaces := 0
	for i, t := range whitespace_sep_tokens {
		if t.IsWhitespace() {
			whitespaces += 1
			cursor = i
		} else {
			break
		}
	}
	return whitespace_sep_tokens[cursor:], whitespaces
}
