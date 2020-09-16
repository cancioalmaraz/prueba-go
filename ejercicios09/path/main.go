package main

import (
	"fmt"
	"os"
)

func main() {
	msg, _ := os.Getwd()
	fmt.Println(msg)
}
