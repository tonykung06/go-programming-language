package main

import (
	"fmt"
	"time"
)

func printNameWithSuffic(name, suffix string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(name, suffix)
	}
}

func main() {
	done := make(chan bool, 2)
	go func() {
		printNameWithSuffic("tony", "!!!", 5)
		done <- true
		time.Sleep(1 * time.Millisecond)
		done <- true
		fmt.Println("DONE!!!")
	}()
	printNameWithSuffic("wini", "!!!", 5)
	<-done
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
