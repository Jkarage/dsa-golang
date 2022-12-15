package main

import "fmt"

func main() {
	char := make(map[rune]int)
	sent := "Ad greendApple"

	for _, v := range sent {
		char[v]++
	}

	for _, v := range sent {
		if char[v] == 1 {
			fmt.Println(string(v))
			return
		}
	}
}
