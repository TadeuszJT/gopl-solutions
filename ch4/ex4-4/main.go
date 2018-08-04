// Write a version of rotate that operates in a single pass.
//
// Tadeusz Tomoszek
// Inverness Public Library 30.7.2018


package main

import (
	"fmt"
)

func main() {
	ints := []int{0, 1, 2, 3, 4, 5}
	rotate(ints, 8)
	fmt.Println(ints)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	n = n % len(s)
	for i := 0; i < len(s)-n; i++ {
		s[i], s[i+n] = s[i+n], s[i]
	}
}


// 0, 1, 2, 3, 4, 5
// swap i0 with i2
// 2, 1, 0, 3, 4, 5 
// swap i1 with i3
// 2, 3, 0, 1, 4, 5
// swap i2 with i4
// 2, 3, 4, 1, 0, 5
// swap i3 with i5
// 2, 3, 4, 5, 0, 1

