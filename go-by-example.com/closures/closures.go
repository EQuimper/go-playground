package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value,
	// which will be updated each time we call nextInt.
	nexInt := intSeq()

	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
