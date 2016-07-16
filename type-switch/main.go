package main

import (
	"go-programming-language/lib"
)

func main() {
	person := lib.Person{"Tony", 25}
	lib.TypeSwitch(person)
}
