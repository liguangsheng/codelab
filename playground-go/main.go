package main

import "fmt"

func main() {
	var text string
	fmt.Scanln(&text)
	fmt.Println(text)
}

func check(eee ...interface{}) {
	for _, e := range eee {
		if err, ok := e.(error); ok && err != nil {
			panic(err)
		}
	}
}

func noop(...interface{}) {}
