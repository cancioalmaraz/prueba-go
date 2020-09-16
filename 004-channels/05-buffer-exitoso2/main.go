package main

import (
	"fmt"
)

var ca chan int = make(chan int, 2)

func main() {
	// Buffered channel (canal con buffer)

	ca <- 42
	ca <- 43

	fmt.Println(len(ca))
	fmt.Println(<-ca)
	fmt.Println(len(ca))
	fmt.Println(<-ca)
	fmt.Println(len(ca))
}
