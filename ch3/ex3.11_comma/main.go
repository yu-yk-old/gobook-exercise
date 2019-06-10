// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "-12345678.0001234"
	s = commaIterate(s)
	fmt.Println(s)
}

func commaIterate(s string) string {
	var buf bytes.Buffer
	n := len(s)
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		n--
		start++
	}

	dotIndex := strings.LastIndexByte(s, '.')

	if dotIndex != -1 {
		n = dotIndex - start
		buf.WriteString(s[start : n%3+start])        // write the first 1/2 digits to the buffer
		for i := n%3 + start; i < dotIndex; i += 3 { // start from the 1/2 digits + start
			if i != start {
				buf.WriteByte(',')
			}
			buf.WriteString(s[i : i+3])
		}
		buf.WriteString(s[dotIndex:])
	} else {
		buf.WriteString(s[start : n%3+start]) // write the first 1/2 digits to the buffer
		for i := n%3 + start; i < n; i += 3 { // start from the 1/2 digits + start
			if i != start {
				buf.WriteByte(',')
			}
			buf.WriteString(s[i : i+3])
		}
	}

	return buf.String()
}
