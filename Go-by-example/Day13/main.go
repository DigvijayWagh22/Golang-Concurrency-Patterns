package main

import "fmt"

// channel syncronization

func worker(done chan bool) {
	fmt.Println("started")
	fmt.Println("working")
	fmt.Println("done")
	done <- true

}
func main() {
	msgch := make(chan bool, 1)
	go worker(msgch)

	<-msgch

}
