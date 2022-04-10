package parser

// handle expressions after all the filtering and whtiespace removal, handle the return codes
// from the parser.checkPattern() function

const unknown_expression = -1 // this is only for dev purposes
const keyword_const = 0
const keyword_var = 1

func handlePattern(cursor *multipleCursor, parser *Parser, code int) {
	switch code {
	case keyword_const:
		handle_const_expression(cursor, parser)
	case keyword_var:
		handle_var_expression(cursor, parser)
	case unknown_expression:
		handle_unknown_expression(cursor, parser) // this is only for testing purposes
	}
}
