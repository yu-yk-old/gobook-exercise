/*
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["Stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["Stdin"])
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// check the key exist or not, if not, initalize the inner map
			if _, ok := counts[fileName]; !ok {
				counts[fileName] = make(map[string]int)
			}
			countLines(f, counts[fileName])
			f.Close()
		}
	}
	for fileName, lineCount := range counts {
		for line, n := range lineCount {
			if n > 1 {
				fmt.Printf("%s: %d\t%s\n", fileName, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
