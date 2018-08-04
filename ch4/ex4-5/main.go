// Write an in-place function to eliminate adjacent duplicates in a []string
// slice.
//
// Tadeusz Tomoszek 31.7.2018

package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "a", "a", "b", "c", "b", "b", "c"}
	s = eliminate(s)
	fmt.Println(s)
}

func eliminate(strings []string) []string {
	j := 0
	var s string
	for i := 0; i < len(strings); i++ {
		if strings[i] != s {
			s = strings[i]
			strings[j] = strings[i]
			j++
		}
	}
	return strings[:j]
}
