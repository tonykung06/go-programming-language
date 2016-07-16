package main

import (
	"fmt"
)

func printNameWithSuffic(name, suffix string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(name, suffix)
	}
}

func main() {
	done := make(chan bool)
	go func() {
		printNameWithSuffic("tony", "!!!", 5)
		done <- true
		done <- true //this goroutine is blocked here as the channel is not buffered and filled with a value already
		fmt.Println("DONE!!!")
	}()
	printNameWithSuffic("wini", "!!!", 5)
	<-done
}
