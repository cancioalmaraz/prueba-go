package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Rutina 1")
	}()
	fmt.Printf("%d\n%d\n", runtime.NumCPU(), runtime.NumGoroutine())

	go func() {
		defer wg.Done()
		fmt.Println("Rutina 2")
	}()
	wg.Wait()
	fmt.Println("Goroutine principal")
}
