package lang

import (
	"fmt"

	"github.com/glaukiol1/gago/lexer"
)

// build a stack output
// very simple for now

func BuildStack(token *lexer.Token, filename string) string {
	return "\n\t Error at " + filename + ":" + fmt.Sprint(token.GetLine())
}
