package lexer

import (
	"fmt"
)

func TestLex() *Lexer {
	downcase_lexed := NewLex("const h = \"hello\"", "<testfile>")
	err := downcase_lexed.Lex(true)
	if err != nil {
		fmt.Println("ERROR", downcase_lexed.errorString)
	}
	return downcase_lexed
}
