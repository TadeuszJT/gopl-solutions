// Rewrite PopCount to use a loop instead of a single expression and run a test to
// compare speeds.
//
// Original : 0.29ns/op
// Loop     : 21.5ns/op
// Much slower!
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
	var i uint
	for ; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//!-
