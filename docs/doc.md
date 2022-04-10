# Gago Examples & Syntax

This folder contains examples and docs for the Gago programming language.

This file contains the basic syntax info. For other docs:

- [Data types](datatypes.md)

## Basic Syntax

The basic syntax for the Gago language.

The Gago programming language is based off of lines, not semicolons (`;`), so expressions must be of one line, except some special expressions: such as `for` `if` ...

### Defining Variables

Defining a variable will add it to the memory map, which can be accessed or manipulated later on.

#### Defining a _constant_ variable

```js
const helloworld = "Hello World"
```

#### Defining a _non-constant_ variable

```js
var helloworld = "Hello World"
```

### Calling functions

To call a function, you would do it like in most languages, but you need to add the `call` keyword in the beginning of the expression.

An example of printing

```js
call print("hello world!")
```

This would print out `hello world!` in the terminal.

`print` accepts as many arguments as you like, they will be printed out seperated by a space.
