package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func(messages chan string) {
		messages <- "Hello"
	}(messages)
	go func(messages chan string) {
		messages <- "Hello2"
	}(messages)

	fmt.Println(<-messages)
	time.Sleep(time.Second)
	fmt.Println(<-messages)

}
