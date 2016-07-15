package main

import (
	"./lib"
)

func main() {
	person := lib.Person{"tony", 25}
	lib.Greet(person, true)
}
