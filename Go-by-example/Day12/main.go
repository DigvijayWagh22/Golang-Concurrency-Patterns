package main

import (
	"fmt"
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
	fmt.Println(<-messages)

	// buffered channels

	msgCh := make(chan string, 2)
	msgCh <- "buffered"
	msgCh <- "channel"

	fmt.Println(<-msgCh)
	fmt.Println(<-msgCh)

}
