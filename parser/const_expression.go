package parser

// handle functions
// where the parser found the KEYWORD_CONST code pattern

// some codes for certain keywords, expressions
const keyword_const = 0
const keyword_var = 1

// base function for handling const expressions
func handle_const_expression(cursor *multipleCursor) {
	cursor.SetIndex(1) // start at index 1
}
