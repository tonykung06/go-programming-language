package main

import (
	"fmt"
)

type Message struct {
	title string
	body  string
}

type Printer func(string)

func Email(message Message, do Printer) {
	_, msg1 := formatMessage(message.title, message.body, "yo")
	do(msg1)
}

func formatMessage(title string, body ...string) (string, string) {
	return title + " -- " + body[0], "Hey!!"
}

func Print(value string) {
	fmt.Print(value)
}

func PrintLine(value string) {
	fmt.Println(value)
}

func CreateCustomPrinter(customMessage string) Printer {
	return func(value string) {
		fmt.Println(value + ", " + customMessage)
	}
}

func main() {
	message := Message{
		title: "my message title",
		body:  "my message body",
	}
	Email(message, CreateCustomPrinter("OMG!!!"))
}
