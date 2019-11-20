package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci_myversion() func() int {
	i := 0
	f0, f1, f2 := 0, 1, 0
	return func() int {
		i += 1
		if i == 1 {
		return f0
		}
		if i == 2 {
		return f1}
		
		f2 = f0 + f1
		f0, f1 = f1, f2

		return f2
	}
}

func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
