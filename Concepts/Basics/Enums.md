### Go "Enums" using `const` and `iota`
The standard Go approach involves creating a new custom integer type and defining related constants within a `const` block, using **`iota`** to automatically generate successive values. 

```go
package main

import "fmt"

// Declare a custom type for type safety
type ServerState int

// Define the possible values as constants
const (
    StateIdle ServerState = iota // iota starts at 0
    StateConnected               // automatically assigned 1
    StateError                   // automatically assigned 2
    StateRetrying                // automatically assigned 3
)

func main() {
    state := StateConnected
    fmt.Println(state) // Prints the integer value, e.g., 1

    // Type safety: a plain int cannot be assigned directly
    // var rawInt int = 2
    // state = rawInt // This would cause a compilation error
    state = ServerState(2) // Explicit casting is allowed
}
```
You can also initialize the first value to `iota+1` to start the index from 1 instead of zero.