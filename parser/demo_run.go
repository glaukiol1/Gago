package parser

import "fmt"

func RunDemoLex() {
	for i, token := range Lex("abcdefghijklmnopqrstuvwxyz") {
		fmt.Println("Token " + fmt.Sprint(i) + "(" + token.value + ") at pos " + fmt.Sprint(token.pos) + ":" + fmt.Sprint(token.line) + " = " + fmt.Sprint(token.code))
	}
}
