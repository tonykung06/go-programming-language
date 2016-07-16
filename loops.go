package main

import (
	"fmt"
)

func basicForLoop() {
	times := 5

	for i := 0; i < times; i++ {
		fmt.Println("basic for loop", i)
	}
}

func whileLoop() {
	i := 0
	times := 5
	for i < times {
		fmt.Println("while loop", i)
		i++
	}
}

func infiniteLoop() {
	i := 0
	times := 5
	for {
		if i >= times {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Println("infinite loop...", i)
		i++
	}
}

type Message struct {
	Title string
	Body  string
}

//used for Array, Slice, String, Map, Channel
func ranges() {
	messages := []Message{
		{"my title", "my message body"},
		{"my title2", "my message body2"},
		{"my title3", "my message body3"},
	}

	for index, messages := range messages {
		fmt.Println("Index:", index, "; message:", messages)
	}

	for index := range messages {
		fmt.Println("Only Index:", index)
	}
}

func main() {
	basicForLoop()
	whileLoop()
	infiniteLoop()
	ranges()
}
