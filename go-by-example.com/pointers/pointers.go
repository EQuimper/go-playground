package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records within your program.

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// The &i syntax gives the memory address of i, i.exercise1.5. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}

// zeroval doesn’t change the i in main,
// but zeroptr does because it has a reference
// to the memory address for that variable.
