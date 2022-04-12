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

#### Assigning a variable to a function result

You can assign functions like any other result.

```js
const name = call input("enter your name: ")
```

the variable name now has the value of whatever input() returned.

### Calling functions

To call a function, you would do it like in most languages, but you need to add the `call` keyword in the beginning of the expression.

An example of printing

```js
call print("hello world!")
```

This would print out `hello world!` in the terminal.

`print` accepts as many arguments as you like, they will be printed out seperated by a space.

An example of taking input from the user

```js
call input("enter your name ")
```

You can also store the result value of a function in a variable, like so:

```js
const name = call input("enter your name ")
```

You can then print the value, like so

```js
const name = call input("enter your name ")
call print("your name is", name)
// examples/name.gago
```

A more compact example of this would be

```js
call print("your name is", call input("enter your name "))
// examples/v0.4name.gago
```

## Math

Math in Gago is like any other language.

```js
const mexpr = 10+20+10
call print("mexpr:", mexpr)
// this will print out 40
```

You can also add together variables, mixing them with numbers too.

```js
// examples/math.gago
const num1 = 10
const num2 = 20
call print("num1 + num2", num1 + num2)
call print("num1 / num2", num1 / num2)
call print("num1 ^ num2", num1 ^ num2)

// output
/*
num1 + num2 30
num1 / num2 0.500000
num1 ^ num2 30
*/
```
