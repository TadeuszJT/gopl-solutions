// Modify the echo program to print os.Args[0] as well.
//
// Solved by Tadeusz Tomoszek 11.7.2018

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
