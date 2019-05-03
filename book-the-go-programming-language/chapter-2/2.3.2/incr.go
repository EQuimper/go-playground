package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	incr(&v) // side effect: v is now 2
	fmt.Println(v)
}
