# Maps

## Some stuff about maps, that I've alreade knew, but always good to remind

The key type is special.
It can only be a comparable type because without the ability to tell if 2 keys are equal,
we have no way to ensure that we are getting the correct value.

## Pointers, copies, et al

An interesting property of maps is that you can modify them without passing
as an address to it (e.g &myMap)

This may make them feel like a "reference type", [but as Dave Cheney describes](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) they are not.

  A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part,
not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic.

Therefore, you should never initialize a nil map variable:
```go
var m map[string]string
```

Instead, you can initialize an empty map or use the make keyword to create a map for you:
```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.
