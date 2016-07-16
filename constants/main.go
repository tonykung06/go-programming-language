package main

import (
	"fmt"
)

const (
	PI       = 3.14
	Language = "Go"
	A        = iota
	B
	C
)

const (
	D = iota
	E
)

func main() {
	fmt.Println(A, B, C, D, E, PI, Language)
}
