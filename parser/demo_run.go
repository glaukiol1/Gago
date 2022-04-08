package parser

import "fmt"

func TestLex() bool {
	downcase_lexed := Lex("abcdefghijklmnopqrstuvwxyz")
	uppercase_lexed := Lex("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	num_lexed := Lex("0123456789")
	passed := true
	for i, token := range downcase_lexed {
		if i == token.code {
			fmt.Println("\033[92m Token :" + token.value + ": OK!\033[00m")
		} else {
			fmt.Println("\033[91m Token :" + token.value + ": FAILED!\033[00m")
			passed = false
		}
	}
	for i, token := range uppercase_lexed {
		if i+26 == token.code {
			fmt.Println("\033[92m Token :" + token.value + ": OK!\033[00m")
		} else {
			fmt.Println("\033[91m Token :" + token.value + ": FAILED!\033[00m")
			passed = false
		}
	}
	for i, token := range num_lexed {
		if i+(26*2) == token.code {
			fmt.Println("\033[92m Token :" + token.value + ": OK!\033[00m")
		} else {
			fmt.Println("\033[91m Token :" + token.value + ": FAILED!\033[00m")
			passed = false
		}
	}
	return passed
}
