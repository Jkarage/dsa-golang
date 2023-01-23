package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("needle", 0, "a value to be searched in the slice")
	flag.Parse()
	h := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	fmt.Println(BinarySearch(h, *n))
}

func BinarySearch(h []int, n int) bool {
	lo := 0
	hi := len(h)
	for lo < hi {
		m := lo + ((hi - lo) / 2)
		if h[m] == n {
			return true
		} else if h[m] < n {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return false

}

// Usage
// go run binary_search.go --needle 3
