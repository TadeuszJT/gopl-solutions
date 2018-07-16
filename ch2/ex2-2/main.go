// Write a program like cf that converts other units such as length and weight.
//
// Solved by Tadeusz Tomoszek 16-7-2018

package main

import (
	"fmt"
	"flag"
	"strconv"
	"os"
	"solutions/ch2/ex2-1"
)

func main() {
	temp   := flag.Bool("t", false, "Convert temperature")
	length := flag.Bool("l", false, "Convert length")
	weight := flag.Bool("w", false, "Convert weight")
	flag.Parse()

	for _, arg := range flag.Args() {
		if x, err := strconv.ParseFloat(arg, 64); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			if *temp {
				c := tempconv.Celsius(x)
				f := tempconv.Fahrenheit(x)
				k := tempconv.Kelvin(x)
				fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
				fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
				fmt.Printf("%s = %s = %s\n", k, tempconv.KToC(k), tempconv.KToF(k))
			}
			if *length {
				fmt.Printf("%gm = %g\"\n", x, x * 39.37)
				fmt.Printf("%g\" = %gm\n", x, x / 39.37)
			}
			if *weight {
				fmt.Printf("%gkg = %glbs\n", x, x*2.20463)
				fmt.Printf("%glbs = %gkg\n", x, x / 2.20463)
			}
		}
	}
}
