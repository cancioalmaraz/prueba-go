package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("SO: %v\nARCH: %v\n", runtime.GOOS, runtime.GOARCH)
}
