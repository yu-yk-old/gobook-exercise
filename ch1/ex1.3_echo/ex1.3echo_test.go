/*
Echo prints its command-line arguments.
Exercise 1.3:
Experiment to measure the difference inrunning time between our potentially inefficient versions
and the one that uses strings.Join. (Section 1.6 illustrates part of the time package,
and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

Benchmark_echo2-8   	 3000000	       395 ns/op	     176 B/op	       8 allocs/op
Benchmark_echo3-8   	20000000	       105 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	gopl/ch1/echo_ex1.3	3.838s

*/
package main

import (
	"strings"
	"testing"
)

var args = []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "a10"}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo3(args []string) {
	// fmt.Println(strings.Join(args[1:], " "))
	strings.Join(args[1:], " ")
}

func Benchmark_echo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func Benchmark_echo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
