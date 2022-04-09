package ast

const AST_TYPE_VAR_DECLARATION = 0   // const or var | first init     | declaration   | assignment
const AST_TYPE_VAR_REDECLARATION = 1 // const or var | not first init | redeclaration | reassignment
