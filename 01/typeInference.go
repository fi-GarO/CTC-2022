package main

import "fmt"

func main() {
	v := 42 // change me!
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

