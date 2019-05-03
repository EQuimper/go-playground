package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println("*p:", *p)
	*p = 2
	fmt.Println("x: ", x)

	s := 10
	t := s
	fmt.Println("t: ", t)
	t = 20
	fmt.Println("s: ", s)
	fmt.Println("t: ", t)
}
