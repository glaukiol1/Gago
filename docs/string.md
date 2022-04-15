# String | Standard Library Module

The string module contains functions for interacting with strings in Gago.

A string is a base data type, so you can create it directly, with no need for a function. Example: `const a = "hello!"`.

## Functions

- concat(str, args...) `concat joins all the arguments together`
- contains(str, substr) `contains searches for the substring in the specified string. Returns a boolean`
- containsAny(str, chars) `containsAny checks if any of the chars are in the set string. Returns a boolean`
- trimSpace(str) `trimSpace removes all trailing and leading whitespaces and returns the new string`
- index(str, substr) `index returns the index of the first instance of substr in string, or -1 if substr is not present in string.`
- len(str) `len returns the length of the string`
- charAt(str, index) `charAt returns the character at the given index, or a IndexError`

## Importing

To import the string module, since its part of the standard library, you can import it by default: `import string`.
