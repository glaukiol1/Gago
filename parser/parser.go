package parser

import (
	"fmt"

	"github.com/glaukiol1/gago/lexer"
)

type Parser struct {
	lexer            *lexer.Lexer     // lexer
	tokens           []*lexer.Token   // tokens
	whitespaceTokens [][]*lexer.Token // tokens seperated by whitesapces
	error            bool             // keep track of errors
	errorString      string           // error string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer, tokens: lexer.GetTokens(), error: false, errorString: ""}
}

func (parser *Parser) sepWhitespaces() {
	var local []*lexer.Token
	for _, t := range parser.tokens {
		if t.IsNewline() {
			parser.whitespaceTokens = append(parser.whitespaceTokens, local)
			local = []*lexer.Token{}
		} else if t.GetCode() == 96 { // EOF
			parser.whitespaceTokens = append(parser.whitespaceTokens, local)
		} else {
			local = append(local, t)
		}
	}
}

func (parser *Parser) Parse(v bool) {
	parser.sepWhitespaces()
	if v {
		for i, whitespace_sep_tokens := range parser.whitespaceTokens {
			fmt.Println("---" + fmt.Sprint(i) + "---")
			for _, v := range whitespace_sep_tokens {
				fmt.Println("token", fmt.Sprint(v.GetValue()), fmt.Sprint(v.GetCode()))
			}
		}
	}
}
