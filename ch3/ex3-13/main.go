// Write const declarations for KB up to YB as compactly as possible.
//
// Solved by Tadeusz Tomoszek 27.7.2018

package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
