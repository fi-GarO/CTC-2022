package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var x, prev1, prev2 int
	return func() int {
		var fib int
		var tmp int
		x += 1
		if x == 1 {
			fib = x - 1
			prev1 = fib
		} else if x == 2 {
			fib = x - 1
		} else {
			fib = prev1 + prev2
		}

		tmp = prev1
		prev1 = fib
		prev2 = tmp

		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
