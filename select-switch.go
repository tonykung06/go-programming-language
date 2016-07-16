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
	c2 := make(chan Person)
	allPeople := People{
		{"Tony", 25},
		{"Wini", 24},
	}

	//call a goroutine that will fill the channel
	go allPeople.putPeopleIntoChannel(c)
	go allPeople.putPeopleIntoChannel(c2)

	for {
		//select the "ready" channel,
		//  if more than one channel is ready, randomly select one of them,
		//  if none is ready, then execute the default case or block if there is no default case
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else { //when the channel is closed, ok is false
				//this could return earlier than the other channels finish
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else { //when the channel is closed, ok is false
				//this could return earlier than the other channels finish
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
