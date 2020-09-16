package main

import (
	_ "./foo"
	"C"
	"fmt"
)

//export Add
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
}
