### Intro
Go contains basic numeric types that can represent sets of either integer or floating-point 
values. There a different types depending on the size of value you require and the 
architecture of the computer where the application is running (e.g. 32-bit and 64-bit). <br>

This concept will concentrate on two number types:
- `int`: e.g. 0, 255, 2147483647. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647). But this will depend on the systems architecture. Most modern computers are 64 bit, therefore int will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

- `float64`: e.g. 0.0, 3.14. Contains the set of all 64-bit floating-point numbers.

For a full list of the available numeric types and more detail see the following resources:
- [Go builtin type declarations](https://github.com/golang/go/blob/master/src/builtin/builtin.go)
- [Digital Ocean - Understanding Data Types in Go](https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go)
- [Arden Labs - Understanding Type in Go](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html)

The Go language provides these basic numeric types:
```bash
# Unsigned Integers
uint8, uint16, uint32, uint64

# Signed Integers
int8, int16, int32, int64

# Real Numbers
float32, float64

# Predeclared Integers
uint, int, uintptr
```

Go supports the standard set of arithmetic operators of `+`, `-`, `*`, `/` and `%` (remainder not modulo).

In Go, assignment of a value between different types requires explicit conversion. For example, to convert an `int` to a `float64` you would need to do the following:
```go
var x int = 42
f := float64(x)

fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
// Output: x is of type: int

fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
// Output: f is of type: float64
```
For more information on Type Conversion please see [A Tour of Go: Type Conversions](https://go.dev/tour/basics/13)

### About Floats
A floating-point number is a number with zero or more digits behind the decimal separator. Examples are -2.4, 0.1, 3.14, 16.984025 and 1024.0. <br>

Different floating-point types can store different numbers of digits after the digit separator - this is referred to as its precision.<br>

Go has two floating-point types:
- `float32`: 32 bits (~6-9 digits precision).
- `float64`: 64 bits (~15-17 digits precision). This is the default floating-point type.
As can be seen, both types can store a different number of digits. This means that trying to store PI in a float32 will only store the first 6 to 9 digits (with the last digit being rounded).<br>

By default, Go will use `float64` for floating-point numbers, unless the floating-point number is:
- assigned to a variable with type `float32`, or
- returned from a function with return type `float32`, or
- passed as an argument to the `float32()` function.