// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func main() {
	data := []string{"one", "one", "two", "two", "two", "one", "three", "three"}
	fmt.Println(delDup(data)) // output: [one two one three]
}

func delDup(s []string) []string {
	slen := len(s)
	for i, j := 0, 0; j < slen; j++ {
		if s[i] == s[j] {
			if i != j {
				copy(s[i:], s[j:slen])
				slen--
				j = i
			}
		} else {
			i++
		}
	}
	return s[:slen]
}
