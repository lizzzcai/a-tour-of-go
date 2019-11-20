# Flow control statements: for, if, else, switch and defer

## For

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```
# For is Go's "while"

The init and post statements are optional.

At that point you can drop the semicolons: C's while is spelled for in Go.

```go
// The init and post statements are optional.
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```


```go
// At that point you can drop the semicolons: C's while is spelled for in Go.
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
## Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
```go
func main() {
	for {
	}
}
```

## If
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
```go
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
```

## If with a short statement
Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.
```go
import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

# If and else
Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow return their results before the call to fmt.Println in main begins.)

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```

## Exercise: Loops and Functions
```go
func Sqrt(x float64) float64 {
	z := x / 2
	for delta := 1.0; delta > 5e-3; delta = math.Abs(x - z*z) {
		z -= (z*z - x) / (2*z)
	}
	return z
}
```

## Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

## Switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
```go
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```
## Switch with no condition
Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.
```go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

## Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

## Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this blog post. https://blog.golang.org/defer-panic-and-recover

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
/* result
counting
done
9
8
7
6
5
4
3
2
1
0
*/
```