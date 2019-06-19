// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

// reverse reverses a slice of ints in place.
func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
