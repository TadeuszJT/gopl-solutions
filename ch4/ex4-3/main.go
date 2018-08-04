// Rewrite reverse to use an array pointer instead of a slice.
//
// Tadeusz Tomoszek
// Inverness Public Library 30.7.2018

package main

import (
	"fmt"
)

func main() {
	ints := [6]int{1, 2, 3, 4, 5, 6}
	reverse2(&ints)
	fmt.Println(ints)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

