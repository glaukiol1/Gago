# Object | Standard Library Module

The object module has functions for interacting with objects in Gago.

## Creating A Object

To create a new object, you run the `object.create()` function. It has no paramaters and returns a blank object.

## Setting a value

To set a value of a object in Gago, you use the `object.set()` function. The syntax is as follows:

```js
object.set(obj, key, value)
```

- `obj` object
- `key` the key that will hold the value (string)
- `value` any value you would like to assign to that key

This function returns a `object` type.

## Getting a value

To get a value, you use the `object.get()` function. The syntax is as follows:

```js
object.get(obj, key)
```

- `obj` object
- `key` the key that you want to get the value of

This returns the value of that key-value pair or `null`.

## Getting all keys

To get a array of all the keys, you use the `object.keys()` function. Its only paramater is the object. It returns a [object](object.md) which contains strings (keys).
