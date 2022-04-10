package ast

import "github.com/glaukiol1/gago/lang"

// literal will be a AST type which will
// hold all literals

type Literal struct {
	AstType int       // the ast type
	Value   lang.Type // the value
}
