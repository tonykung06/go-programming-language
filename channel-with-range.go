package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type People []Person

func (people People) putPeopleIntoChannel(c chan Person) {
	for _, s := range people {
		c <- s
	}
	close(c)
}

func main() {
	c := make(chan Person)
	allPeople := People{
		{"Tony", 25},
		{"Wini", 24},
	}

	//call a goroutine that will fill the channel
	go allPeople.putPeopleIntoChannel(c)

	for p := range c {
		fmt.Println(p.Name)
	}
}
