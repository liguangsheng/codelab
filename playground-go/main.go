package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := Int64Heap{4, 2, 5, 5, 2, 1, 3, 3}
	heap.Init(&nums)
	fmt.Println(nums)

	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
	heap.Push(&nums, int64(1))
	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
	fmt.Println(heap.Pop(&nums))
}

func check(eee ...interface{}) {
	for _, e := range eee {
		if err, ok := e.(error); ok && err != nil {
			panic(err)
		}
	}
}

func noop(...interface{}) {}
