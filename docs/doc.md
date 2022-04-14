# Gago Examples & Syntax

This folder contains examples and docs for the Gago programming language.

This file contains the basic syntax info. For other docs:

- [Data types](datatypes.md)
- [Array Module](array.md)

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

#### Reassigning Variables

To reassign a variable, you must make sure that you are reassigning the same type. If you try to assign a `int` to a `string`, you would get a error that looks like this:

```md
TypeError: Can not assign value of type `int` to variable of type `string`
```

If you are sure that the types match, you can eassign a variable via the `reset` keyword. An example script would be:

```js
var test = "test variable"
call print("test:", test)
reset test = call input("new value for test: ")
call print("test:", test)
// examples/reassignment.gago
```

This would get the user input and assign it to the variable `test`.

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
call print("num1 ** num2", num1 ^ num2)

// output
/*
num1 + num2 30
num1 / num2 0.500000
num1 ** num2 -9223372036854775808
*/
```

Now you might ask, why is the output of `num1 ** num2` (num1 raised to the power of num2) `-9223372036854775808`. Well, it is because the actual value of `num1 ** num2` is more than the max value of `int64`. This is called a integer overflow, which causes the value to wrap and become negative.

You can also use math with the `float` datatype. A simple example would be

```js
// examples/float.gago
const fl = 10.123456
call print("fl =", fl)
call print("fl+10 =", fl+10)
call print("fl*2 =", fl*2)
call print("fl**2 =", fl**2)

// output
/*
fl = 10.123456
fl+10 = 20.123456
fl*2 = 20.246912
fl**2 = 102.484361
*/
```

In Gago, floats have a percision of 6. Which means that there can be 6 integers after the decimal point. If there are more than 6, it is rounded.

## Builtins

Docs for builtin functions

### `print`

Print `args` seperated by a space. There can be as many args as you want. They will be printed to standard out.
They can be of any type.

### `input`

To get input from a user, you run the `input(message)` function, where `message` is a string that will be printed to standard out in the line where the input starts.

Example:

```js
// examples/name.gago
const name = call input("whats your name? ")
call print("your name is", name)
```

Now you can enter your name in the command line, and it would print out `your name is <yourname>`.

### `exit`

To exit from the process, you can run `exit()`, with an optional `<int code>` argument. The `code` argument must be of type int.

Example:

```js
exit(0)
```
