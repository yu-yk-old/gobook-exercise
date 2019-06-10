// Exercise 3.10: Write a non-recursive version of comma,
// using bytes.Buffer instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "12345678"
	s = commaIterate(s)
	fmt.Println(s)
}
func commaIterate(s string) string {
	var buf bytes.Buffer
	n := len(s)
	buf.WriteString(s[:(n % 3)])    // write the first 1/2 digits to the buffer
	for i := n % 3; i < n; i += 3 { // start from the 1/2/3 digits
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
