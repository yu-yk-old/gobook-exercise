// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	const usage = `
	Error: wrong numbers of args
	usage: go run main.go (str1) (str2)`

	if len(os.Args) != 3 {
		log.Fatal(usage)
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if anagrams(s1, s2) {
		fmt.Printf("\"%s\" and \"%s\" are anagrams\n", s1, s2)
	} else {
		fmt.Printf("\"%s\" and \"%s\" are not anagrams\n", s1, s2)
	}
}

func anagrams(s1, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}
	i := uint32(0)
	for _, c := range s1 {
		i ^= uint32(c)
	}

	for _, c := range s2 {
		i ^= uint32(c)
	}

	if i == 0 {
		return true
	}

	return false
}
