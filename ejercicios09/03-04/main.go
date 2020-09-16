package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mx sync.Mutex
var wg sync.WaitGroup

func main() {
	var incremento int64
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&incremento, 1)
			fmt.Println(atomic.LoadInt64(&incremento))
		}()
	}
	wg.Wait()
	fmt.Println(incremento)

}
