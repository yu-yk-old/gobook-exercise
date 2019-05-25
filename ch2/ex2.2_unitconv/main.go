// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf
// that reads numbers from its command-line arguments or from the standard input if
// there are no arguments, and converts each number into units like temperature in
// Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl/ch2/ex2.2_unitconv/unitconv"
)

var unit = flag.String("u", "t", "unit: (t|l|w) for temperature, length, weight")

func main() {
	flag.Parse()

	if len(os.Args) < 4 {
		// ask for stdin
		input := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pls input a number: ")
		for input.Scan() {
			value, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			conv(unit, value)
			fmt.Printf("Pls input a number: ")
		}
	}

	for _, arg := range os.Args[3:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		conv(unit, value)
	}

}

func conv(unit *string, value float64) {
	switch *unit {
	case "t":
		f := unitconv.Fahrenheit(value)
		c := unitconv.Celsius(value)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToC(f), c, unitconv.CToF(c))
	case "l":
		f := unitconv.Feet(value)
		m := unitconv.Meter(value)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FtoM(f), m, unitconv.MtoF(m))
	case "w":
		p := unitconv.Pound(value)
		kg := unitconv.Kilogram(value)
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PtoKG(p), kg, unitconv.KGtoP(kg))
	}
}
