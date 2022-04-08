package parser

func Lex(toparse string) []*Token {
	var tokens []*Token
	for pos, char := range toparse {
		tokens = append(tokens, NewToken(string(char), pos, 1))
	}
	return tokens
}
