package main

import (
	"fmt"
)

type Message struct {
	title string
	body  string
}

func Email(message Message) {
	_, msg1 := formatMessage(message.title, message.body, "yo")
	fmt.Println(msg1)
}

func formatMessage(title string, body ...string) (string, string) {
	return title + " -- " + body[0], "Hey!!"
}

func main() {
	message := Message{
		title: "my message title",
		body:  "my message body",
	}
	Email(message)
}
