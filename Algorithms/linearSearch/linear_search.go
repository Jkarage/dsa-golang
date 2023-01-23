package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("needle", 0, "a value being searched")
	flag.Parse()
	h := []int{1, 3, 5, 7, 8, 9, 10}
	fmt.Println(LinearSearch(h, *n))
}

func LinearSearch(h []int, n int) bool {
	for _, v := range h {
		if n == v {
			return true
		}
	}
	return false
}

// Usage
// go run linear_search.go --needle 9
