// Write and in-place function that squashes each run of adjacent Unicode
// spaces in a UTF-8 encoded []byte slice into a single ASCII space
//
// Tadeusz Tomoszek 

package main

import (
	"fmt"
)

func squash(s string) {
	for _, r := range s {
		fmt.Println(string(r))
	}
}

func main() {
	squash("Hello, 世界")
}
