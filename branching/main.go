package main

import (
	"fmt"
	person2 "github.com/go-programming-language/branching/person"
	"github.com/go-programming-language/lib"
	_ "gopkg.in/mgo.v2"
)

func main() {
	person := lib.Person{"Tony", 25}
	lib.Greet(person, true)
	privatePerson := person2.PrivatePerson{
		PrivateName: "TOMMY",
	}
	fmt.Println("DONE", privatePerson)
}
