package main

import (
	"fmt"
)

func main() {
	var s []int
	s = make([]int, 3, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s = append(s, []int{4, 5, 6}...)

	fmt.Println(s[:])
}
