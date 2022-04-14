package parser

// handle expressions after all the filtering and whtiespace removal, handle the return codes
// from the parser.checkPattern() function

const unknown_expression = -1 // this is only for dev purposes
const keyword_const = 0
const keyword_var = 1
const keyword_call = 2
const keyword_reset = 3
const keyword_import = 4

func handlePattern(cursor *multipleCursor, parser *Parser, code int) {
	switch code {
	case keyword_const:
		handle_const_expression(cursor, parser)
	case keyword_var:
		handle_var_expression(cursor, parser)
	case keyword_call:
		handle_call_expression(cursor, parser)
	case keyword_reset:
		handle_reset_expression(cursor, parser)
	case keyword_import:
		handle_import_expression(cursor, parser)
	case unknown_expression:
		handle_unknown_expression(cursor, parser) // this is only for testing purposes
	}
}
