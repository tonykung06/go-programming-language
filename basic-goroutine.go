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
	go printNameWithSuffic("tony", "!!!", 5)
	printNameWithSuffic("wini", "!!!", 5)
	time.Sleep(100 * time.Millisecond)
}
