package main

import (
	"fmt"
	person2 "github.com/go-programming-language/branching/person"
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
	privatePerson := person2.PrivatePerson{
		PrivateName: "TOMMY",
	}
	fmt.Println("DONE", privatePerson)
}
