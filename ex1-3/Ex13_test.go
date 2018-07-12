// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.
//
// Solved by Tadeusz Tomoszek 11.7.2018

package main

import (
	"strings"
	"testing"
	"os"
)

// 12.9ns/op
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
}

// 13.7ns/op
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
}

// 4.53ns/op
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}

