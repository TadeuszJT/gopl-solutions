// Write a version of PopCount that uses the expression x&(x-1) which clears
// the right-most bit of x.
//
// 34.3ns/op - kinda slow.
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
	for ; x > 0; x = x&(x-1) {
		count++
	}
	return count
}

//!-
