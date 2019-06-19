// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	bitDiff := bitDiff(c1, c2)
	fmt.Printf("bit differences = %d\n", bitDiff)
}

func bitDiff(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		diff := c1[i] ^ c2[i]
		for diff != 0 {
			diff &= (diff - 1)
			count++
		}
	}
	return count
}
