package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\nArquitectura: %s\n", runtime.GOOS, runtime.GOARCH)
}
