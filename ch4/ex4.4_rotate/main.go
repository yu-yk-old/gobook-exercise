// Exercise 4.4: Write a version of rotate that operates in a single pass.
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s, 2, 'l')
	fmt.Println(s) // "[2 3 4 5 0 1]"

	rotate(s, 4, 'r')
	fmt.Println(s)
}

func rotate(s []int, offset int, direction rune) {
	slen := len(s)
	temp := make([]int, offset)
	switch direction {
	case 'l':
		copy(temp, s[:offset])
		copy(s[:slen-offset], s[offset:])
		copy(s[slen-offset:], temp)
	case 'r':
		copy(temp, s[slen-offset:])
		copy(s[offset:], s[:slen-offset])
		copy(s[:offset], temp)
	}
}
