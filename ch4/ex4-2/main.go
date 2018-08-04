// Write a program to print hashes of stdin with flags to select which type
//
// Tadeusz Tomoszek 28.7.2018

package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	f1 := flag.Bool("sha384", false, "Print sha384 version of Stdin")
	flag.Parse()

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errorred\n")
		return
	}

	if *f1 != false {
		c := sha512.Sum384(buf)
		fmt.Println("Printing sha384:")
		fmt.Println(c)
	} else {
		c := sha256.Sum256(buf)
		fmt.Println("Printing sha256:")
		fmt.Println(c)
	}
}
