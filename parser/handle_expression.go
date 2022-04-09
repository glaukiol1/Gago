package parser

// handle expressions after all the filtering and whtiespace removal, handle the return codes
// from the parser.checkPattern() function

func handlePattern(cursor *multipleCursor, parser *Parser, code int) {
	switch code {
	case keyword_const:
		handle_const_expression(cursor, parser.lexer, parser)
	}
}
