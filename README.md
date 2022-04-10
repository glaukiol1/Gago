# Gago | Programming Language Built in Go

_if you are looking for the docs, go [here](docs/doc.md)_

Gago is a _interpreted_ programming language. It is fully written in Go.

Gago includes:

- Lexer
- Parser
- VM
- Module implementation
- Easy embedding into your Go application

The gago standard library is not yet written. (as of v0.2-alpha)

The gago builtins module: (as of v0.2-alpha)

- `print(args...)` type: function. Doc: `prints the specified values seperated by a space`
- `teststring` type: global variable. Use: `testing purposes`

## Install

Download from the [releases page](https://github.com/glaukiol1/gago/releases)

## Objectives

Gago started as an experiment to find out how hard it would be to build a programming language in Go. With all the standard library written in Go, it would be aster than many programming languages, such as Python.

It can also be used as a wrapper around Go, since any Gago expression can be written in Go.

## Status

_As of `v0.2-alpha`_

gago currently:

- Lexes all tokens correctly
- Parses some simple expressions to AST
- Is able to run the AST available right now (in the VM)
- Has a builtin module, which will be extended.

Taking the project further is the main goal. Making a simple language which can run just like any programming language would be a end goal. Once Gago is ready to run some of the more basic code, running benchmarks will be added.

## License

This project is licensed under the MIT licence.
