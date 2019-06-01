// goos: darwin
// goarch: amd64
// pkg: gopl/ch2/ex2.3_popcount/popcount
// BenchmarkPopCount-8    	2000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCount2-8   	100000000	        16.0 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCount3-8   	100000000	        11.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCount4-8   	300000000	         4.88 ns/op	       0 B/op	       0 allocs/op
// PASS
// coverage: 100.0% of statements
// ok  	gopl/ch2/ex2.3_popcount/popcount	5.534s
// Success: Benchmarks passed.

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
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 Exercise 2.3:
// Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.
func PopCount2(x uint64) int {
	var count int
	for i := uint8(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// PopCount3 Exercise 2.4:
// Write a version of PopCount that counts bits by shifting its argument through 64 bit positions,
// esting the rightmost bit each time. Compare its performance to the table-lookup version.
func PopCount3(x uint64) int {
	var count int
	for x != 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

// PopCount4 Exercise 2.5:
// The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact, and assess its performance.
func PopCount4(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
