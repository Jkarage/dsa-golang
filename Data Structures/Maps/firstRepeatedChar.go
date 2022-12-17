package main

import "fmt"

func main() {
	sent := "green apple"
	runesMap := make(map[rune]int)
	fmt.Println(repeatedChar(sent, runesMap))
}

func repeatedChar(s string, m map[rune]int) string {
	for _, r := range s {
		m[r]++
	}
	for _, r := range s {
		if m[r] > 1 {
			return string(r)
		}
	}
	return ""
}
