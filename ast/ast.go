package ast

const AST_TYPE_VAR_DECLARATION = 0   // const or var | first init     | declaration   | assignment
const AST_TYPE_VAR_REDECLARATION = 1 // const or var | not first init | redeclaration | reassignment
const AST_TYPE_VARIABLE_ACCESS = 2   // variable_access.go
const AST_TYPE_LITERAL = 3           // literal.go
const AST_TYPE_FUNC_CALL = 4         // call a function | methodcall.go
const AST_TYPE_MATH_EXPR = 5         // math expression | mathexpr.go
const AST_TYPE_IMPORT = 6            // import module | import.go
