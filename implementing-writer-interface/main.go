package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (person *Person) Rename(newName string) {
	person.Name = newName
}

func (person *Person) Write(p []byte) (n int, err error) {
	newName := string(p)
	person.Rename(newName)

	n = len(newName)
	err = nil
	return
}

func main() {
	person1 := Person{
		Name: "Tony",
	}
	fmt.Fprintf(&person1, "The length is %d", 10)
	fmt.Println(person1)
}
