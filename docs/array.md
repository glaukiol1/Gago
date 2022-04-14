# Array | Standard Library Module

The Array module contains all methods for interacting with arrays in Gago.

## Creating A Array

To create a new array, you use the `array.create()` function. This function takes as many paramaters as you like, and it returns a new `slice` type.

```js
const myarray = call array.create(1, "Hello!", true, false, "World!")
```

Arrays in Gago can contain any datatype.

## Accessing a element

Arrays in Gago are 0-based, which means the first index starts at 0. An example of indexes would be: `0,1,2,3,4`. This represents 5 values.

To access a element you call the `array.access()` function. Its syntax is like this:

```js
array.access(slice, index)
```

`slice` would be your array, and `index` would be the 0-based index in which the element you want to access is at. It throws a `IndexError` if `index` is out of bounds.

## Getting the length

To get the length of an array, you use the `array.len()` method. You pass an argument of type slice to it, which is your array. It returns the length of that array. __NOTE__: It returns a 1-based index, so if you want to get the last element of an array, you would do something like this: `array.len(myarray)-1`!

## Removing elements

Gago has two functions to help with removing elements, `pop` and `shift`.

### `pop`

Pop _pops_ the last element of the array. It directly modifies the array, and returns `null`.

### `shift`

Shift _shifts_ the array one element; it removes the first element. It also directly modifies the array, and returns `null`.

## Slicing the slice

To _slice the slice_, or to get a subslice, you call the `array.slice()` function. Its syntax is as follows:

```js
array.slice(slice, startindx, endindx)
```

- `slice` is your array
- `startindx` is the index you want to start slicing at (0-based).
- `endindx` is the index you want to stop slicing at. (0-based). Set this to `-1` if you want to get all elements.
