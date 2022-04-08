package lexer

import (
	"fmt"
)

func TestLex() *Lexer {
	downcase_lexed := NewLex("abcdefghijklmnopqrstuvwxyz\\nABCDEFGHIJKLMNOPQRSTUVWXYZ\\n0123456789\\n[]{};:/\\\"'.,<>?|=+-_!@#$%^&*~`() ", "<testfile>")
	err := downcase_lexed.Lex(true)
	if err != nil {
		fmt.Println(downcase_lexed.errorString)
	}
	return downcase_lexed
}
