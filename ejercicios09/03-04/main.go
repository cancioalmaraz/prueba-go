package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mx sync.Mutex
var wg sync.WaitGroup

func main() {
	var incremento int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mx.Lock()
			nv := incremento
			runtime.Gosched()
			nv++
			incremento = nv
			mx.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(incremento)

}
