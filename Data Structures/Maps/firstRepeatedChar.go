package main

import (
	"fmt"
)

func main() {
	sent := "green apple"
	runesSet := make(map[rune]bool)
	fmt.Println(repeatedChar(sent, runesSet))
}

func repeatedChar(s string, m map[rune]bool) string {
	for _, r := range s {
		if b, ok := m[rune(r)]; !ok || !b {
			m[rune(r)] = true
		} else if ok {
			return string(rune(r))
		}
	}
	return ""
}
