// Write a function that reports whether two strings are anagrams of each other.
//
// Solved by Tadeusz Tomoszek 27.7.2018

package main

import (
	"fmt"
	"strings"
)

func anagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	m := make(map[rune]int)
	for _, r := range a {
		m[r]++
	}
	for _, r := range b {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(anagram("BENIS", "sineB"))
}
