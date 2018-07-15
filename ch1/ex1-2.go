// Modify echo to print the index and value of each arg, one per line.
//
// Solved by Tadeusz Tomoszek 11.7.2018

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Println("index:", i, "value:", v)
	}
}
