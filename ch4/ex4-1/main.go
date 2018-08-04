// Write a function that counts the number of bits that are different in two
// sha256 hashes
//
// Solved by Tadeusz Tomoszek 28.7.2018


package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	c1 := sha256.Sum256([]byte("Tadeusz"))
	c2 := sha256.Sum256([]byte("Benis"))
	fmt.Println(Benis(c1, c2))
}

var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Benis(a, b [32]byte) int {
	num := 0
	for i := 0; i < 32; i++ {
		xor := a[i] ^ b[i]  // Highlights different bits
		num += int(pc[xor]) // Counts bits
	}
	return num
}



