package parser

import "fmt"

func RunDemoLex() {
	downcase_lexed := Lex("abcdefghijklmnopqrstuvwxyz")
	uppercase_lexed := Lex("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i, token := range downcase_lexed {
		if i == token.code {
			fmt.Println("\033[92m Token :" + token.value + ": OK!\033[00m")
		} else {
			fmt.Println("\033[91m Token :" + token.value + ": FAILED!\033[00m")
		}
	}
	for i, token := range uppercase_lexed {
		if i+26 == token.code {
			fmt.Println("\033[92m Token :" + token.value + ": OK!\033[00m")
		} else {
			fmt.Println("\033[91m Token :" + token.value + ": FAILED!\033[00m")
		}
	}
}
