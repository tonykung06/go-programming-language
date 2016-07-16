package main

import (
	"fmt"
)

type Message struct {
	title string
	body  string
	body2 string
}

func main() {
	message := Message{
		title: "my message title",
		body:  "my message body",
	}
	message.body2 = "my message body 2"

	fmt.Println(message.body2)
}
