package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel (canal sin buffer)
	ca := make(chan int)

	go func() {
		ca <- 42
		ca <- 43
	}()

	go func() {
		fmt.Println(<-ca)
		fmt.Println(<-ca)
	}()
	time.Sleep(2 * time.Second)
}

func foo() {

}

func bar() {

}
