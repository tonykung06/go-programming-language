package lib

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Greet(person Person, showAlternative bool) {
	if suffix := "!!!"; showAlternative {
		fmt.Println("Hey", person.Name, suffix)
	} else {
		fmt.Println("Hello", person.Name, suffix)
	}
}
