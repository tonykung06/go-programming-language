package main

import (
	"github.com/go-programming-language/lib"
)

func main() {
	person := lib.Person{"Tony", 25}
	lib.Greet(person, true)
}
