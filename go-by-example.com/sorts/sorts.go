package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"d", "b", "e", "a", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{10, 2, 5, 3, 1, 9, 7}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}
