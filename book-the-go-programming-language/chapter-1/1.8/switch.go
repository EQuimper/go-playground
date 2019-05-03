package main

import "fmt"

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}

func main() {
	res := Signum(10)
	fmt.Println(res)
	res2 := Signum(-10)
	fmt.Println(res2)
}
