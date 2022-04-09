package lexer

import (
	"fmt"
)

func TestLex() *Lexer {
	downcase_lexed := NewLex("const h = 'hello'", "<testfile>", true)
	err := downcase_lexed.Lex()
	if err != nil {
		fmt.Println("ERROR", downcase_lexed.errorString)
	}
	return downcase_lexed
}
