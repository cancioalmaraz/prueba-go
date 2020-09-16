package main

import "fmt"

func main() {
	// Buffered channel (canal con buffer)
	ca := make(chan int, 1)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
}
