package parser

import "github.com/glaukiol1/gago/lexer"

// handle expressions after all the filtering and whtiespace removal, handle the return codes
// from the parser.checkPattern() function

func handlePattern(cursor *multipleCursor, lexer *lexer.Lexer, code int) {
	switch code {
	case keyword_const:
		handle_const_expression(cursor, lexer)
	}
}
