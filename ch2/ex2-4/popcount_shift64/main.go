// Modify PopCount to count bits by shifting through all 64 bit positions in an
// unsigned integer.
//
// 111ns/op - Really slow.
//
// Solved by Tadeusz Tomoszek 16.7.2018

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		count += int((x >> i)&1)
	}
	return count
}

//!-
