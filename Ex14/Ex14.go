// Modify dup2 to print all the names of all files in which each duplicated line
// occurs.
//
// Solved by Tadeusz Tomoszek 11.7.2018

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			text := input.Text()
			if text == "" { // skip lines with nothing
				continue
			}
			if counts[text] == nil {
				counts[text] = make(map[string]int)
			}
			counts[text][arg]++
		}
		f.Close()
	}

	for line, files := range counts {
		total := 0
		for  _, count := range files {
			total += count
		}
		if total > 1 {
			for file, count := range files {
				fmt.Printf("%v:%d\t", file, count)
			}
			fmt.Println(line)
		}
	}
}
