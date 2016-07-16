package main

import (
	"fmt"
)

type Renamble interface {
	Rename(newName string)
}

type Person struct {
	Name string
}

func (person *Person) Rename(newName string) {
	person.Name = newName
}

func RenameToFrog(r Renamble) {
	r.Rename("Frog")
}

func main() {
	person1 := Person{
		Name: "Tony",
	}
	RenameToFrog(&person1)
	fmt.Println(person1)
}
