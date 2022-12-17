package main

import "fmt"

func main() {
	employeeNumber := map[string]int{}
	idNumber := map[string]int{}
	fmt.Println(equal(employeeNumber, idNumber))
}

// Maps are not comparable like slices. But we can Iterate over maps
// And lookup of key and value equality
// Returns true If two different maps are equal in terms of key value pairs
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
