package lexer

import (
	"fmt"
)

func TestLex(v bool) *Lexer {
	downcase_lexed := NewLex("const h = 'hello'\nvar h1 = 'hello1'\nteststring\ncall print('h', h, 'h1', h1, 'teststring', teststring)", "<testfile>", v)
	err := downcase_lexed.Lex()
	if err != nil {
		fmt.Println("ERROR", downcase_lexed.errorString)
	}
	return downcase_lexed
}
