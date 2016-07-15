package lib

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Greet(person Person, showAlternative bool) {
	if suffix := getSuffix(person.Name); showAlternative {
		fmt.Println("Hey", person.Name, suffix)
	} else {
		fmt.Println("Hello", person.Name, suffix)
	}
}

func getSuffix(name string) (suffix string) {
	switch name {
	case "Tony", "tony":
		suffix = "???"
		fallthrough
	case "wini", "Wini":
		suffix = "----"
	default:
		suffix = "!!!"
	}

	switch {
	case name == "Tony", name == "tony":
		suffix = "~~~~~~~~~~"
	}
	return
}

func TypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Person:
		fmt.Println("Person")
	default:
		fmt.Println("unknown")
	}
}
