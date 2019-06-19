// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var sha = flag.Uint("SHA", 256, "SHA option: 256|384|512")
var usage = "Usage: main.go [-SHA (256|384|512)] input\n"

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	input := flag.Arg(0)
	switch *sha {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
	}

}
