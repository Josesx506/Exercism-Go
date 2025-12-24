### Strings
A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters. <br>

A string literal is defined between double quotes:
```go
const name = "Jane"
```
Strings can be concatenated via the `+` operator:
```go
"Jane" + " " + "Austen"
// => "Jane Austen"
```
Some special characters need to be escaped with a leading backslash, such as `\t` for a tab and `\n` for a new line in strings.
```go
"How is the weather today?\nIt's sunny"  
// =>
// How is the weather today?
// It's sunny
```
The **strings** package contains many useful functions to work on strings. For more 
information about string functions, check out the [strings package documentation](https://pkg.go.dev/strings). Here are some examples:
```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"
```

### Special characters
As mentioned previously, some special characters need to be escaped with a leading backslash. Below is a more detailed list of what those special characters are:

| Value | Description |
| :--- | :--- |
| \a | Alert or bell |
| \b | Backspace |
| \\ | Backslash |
| \t | Horizontal tab |
| \n | Line feed or newline |
| \f | Form feed |
| \r | Carriage return |
| \v | Vertical tab |
| \' | Single quote |
| \" | Double quote |

```go
fmt.Println("Joe\nWilliam\nJack\nAverell") 
// Output:
// Joe
// William
// Jack
// Averell
```

### Raw string literals
Raw string literals use backticks (***`***) as their delimiter instead of double quotes and 
all characters in it are interpreted literally, meaning that there is no need to escape 
characters or newlines. <br>
This is an example of a multiline string:
```go
const daltons = `Joe
William
Jack
Averell`
```
This string has multiple lines. More specifically, it has 3 \n. However, because we are a 
raw string literal, we don't need to put the \n explicitly in the string. The newlines will 
be interpreted literally. <br>

Here are some other examples comparing raw string literals with regular string literals:
```go
`abc`
// same as "abc"

"\"" // regular string literal with 1 character: a double quote
// same as `"` a raw string literal with 1 character: a double quote

`\n
` // raw string literal with 3 character: a backslash, an 'n', and a newline
// same as "\\n\n"  a regular string literal with 3 characters: a backslash, an 'n', and a newline

"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'
```

### Stringers
Stringer is an interface for defining the string format/representation of values. This is 
similar to `__repr__` in python methods<br>
The interface consists of a single String method:
```go
type Stringer interface {
    String() string
}
```
Types that want to implement this interface must have a `String()` method that returns a 
human-friendly string representation of the type. The `fmt` package (and many others) will 
look for this method to format and print values.

Example: Distances <br>
Assume we are working on an application that deals with geographical distances measured in different units. We have defined types DistanceUnit and Distance as follows:
```go
type DistanceUnit int

// Mapped as indexes for the stringer method
const (
    Kilometer    DistanceUnit = 0
    Mile         DistanceUnit = 1
)
 
type Distance struct {
    number float64
    unit   DistanceUnit
}
```
In the example above, *Kilometer* and *Mile* are constants of type *DistanceUnit*. <br>

These types do not implement interface Stringer as they lack the String method. Hence fmt functions will print Distance values using Go's "default format":
```go
mileUnit := Mile
fmt.Sprint(mileUnit)
// => 1
// The result is '1' because that is the underlying value of the 'Mile' constant (see constant declarations above) 

dist := Distance{number: 790.7, unit: Kilometer}
fmt.Sprint(dist)
// => {790.7 0}
// not a very useful output!
```
In order to make the output more informative, we implement interface Stringer for 
`DistanceUnit` and `Distance` types by adding a String method to each type:
```go
func (sc DistanceUnit) String() string {
    units := []string{"km", "mi"}
    return units[sc]
}

func (d Distance) String() string {
    return fmt.Sprintf("%v %v", d.number, d.unit)
}
```
`fmt` package functions will call these methods when formatting Distance values:
```go
kmUnit := Kilometer
kmUnit.String()
// => km

mileUnit := Mile
mileUnit.String()
// => mi

dist := Distance{
    number: 790.7,
    unit: Kilometer,
}
dist.String()
// => 790.7 km
```