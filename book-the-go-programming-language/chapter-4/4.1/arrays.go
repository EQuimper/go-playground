// array are fixed-length aggregation
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [3]int{}

	fmt.Println(len(a))
	fmt.Println(a[0])

	// this will let us not having to declare the length, this will take the
	// amount of number in the init
	p := [...]string{"a", "b", "c", "d"}
	fmt.Println(len(p))

	symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{99: -1}
	fmt.Println(len(r))
}
