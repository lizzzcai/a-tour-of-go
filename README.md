# a_tour_of_go


## import package

```go
package main

import (
    "fmt"
    "math"
	"math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    // In Go, a name is exported if it begins with a capital letter.
    fmt.Println(math.Pi)

}
```

## Functions

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
// or
func add(x, y int) int {
	return x + y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values, "naked" return. used only in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
    fmt.Println(add(42, 13))
    // multiple results
	a, b := swap("hello", "world")
    fmt.Println(a, b)
    
}

```

## To init the variables

```go
package main
//package level
var c, python, java bool
// A var declaration can include initializers, one per variable.
var i, j int = 1, 2

func main() {
    // function level
    // Variables declared without an explicit initial value are given their zero value.
    var i int
	var f float64
	var b bool
    var s string
    
    var i, j int = 1, 2
    // Short variable declarations, type inference
    k := 3
	c, python, java := true, false, "no!"
    fmt.Println(i, j, c, python, java)
}

```

## Basic types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

## Type conversions / Type inference
in Go assignment between items of different type requires an explicit conversion.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
// or
i := 42
f := float64(i)
u := uint(f)

var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

var v = i // int
```

## Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

An untyped constant takes the type needed by its context.

```go
const Pi = 3.14 // An untyped constant takes the type needed by its context.
const pi int = 3 

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

```

