package main

import (
	"fmt"
	"os"

	"gopl/ch2/ex2.3_popcount/popcount"
)

func main() {
	a, b, c, d := uint64(158), uint64(158), uint64(158), uint64(158)
	fmt.Fprintf(os.Stdout, "%b\n", a)
	fmt.Fprintf(os.Stdout, "%b\n", b)
	fmt.Fprintf(os.Stdout, "%b\n", c)
	fmt.Fprintf(os.Stdout, "%b\n", d)
	w := popcount.PopCount(a)
	x := popcount.PopCount2(b)
	y := popcount.PopCount3(c)
	z := popcount.PopCount4(d)
	fmt.Fprintln(os.Stdout, w)
	fmt.Fprintln(os.Stdout, x)
	fmt.Fprintln(os.Stdout, y)
	fmt.Fprintln(os.Stdout, z)
}
