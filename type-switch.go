package main

import (
	"./lib"
)

func main() {
	person := lib.Person{"Tony", 25}
	lib.TypeSwitch(person)
}
