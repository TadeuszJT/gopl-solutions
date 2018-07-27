// Write a non-recursive version of comma using bytes.Buffer
//
// Tadeusz Tomoszek 27.7.2018

package main

import(
	"fmt"
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer
	off := len(s) % 3
	for i := 0; i < len(s); i++ {
		if (i-off) % 3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("2342612234261212123741234"))
}

