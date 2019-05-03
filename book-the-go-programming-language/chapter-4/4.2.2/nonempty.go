package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"one", "", "three", "", "five", "six", "seven", "eight"}
	fmt.Printf("%q\n", nonempty(data))

	ints := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(remove(ints, 2))
}
