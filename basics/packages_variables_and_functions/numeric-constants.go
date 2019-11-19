/*
Numeric Constants
Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

Try printing needInt(Big) too.

(An int can store at maximum a 64-bit integer, and sometimes less.)
*/

package main

import "fmt"

const (
	// Create a huge number of shifting a 1 bite left 100 places
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {return x*10 + 1}
func needFloat(x float64) float64 {return x * 0.1}

func main() {
	//fmt.Printf("type: %T, %v\n", Big, Big)
	fmt.Printf("type: %T, %v\n", Small, Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// An untyped constant takes the type needed by its context.
	const pi = 3
	fmt.Println(needFloat(pi))

	// below is not working
	// var pi = 3
	// fmt.Println(needFloat(pi))
	// error: cannot use pi (type int) as type float64 in argument to needFloat
}