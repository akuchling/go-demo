// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func runetype(r rune) string {
   if unicode.IsDigit(r) {
      return "digit"
   } else if unicode.IsLower(r) {
      return "lowercase letter"
   } else if unicode.IsUpper(r) {
      return "uppercase letter"
   } else if unicode.IsSpace(r) {
      return "whitespace"
   } else if unicode.IsPunct(r) {
      return "punctuation"
   } else {
      return "unknown"
   }
}

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	categorycount := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		categorycount[runetype(r)]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for s, n := range categorycount {
		fmt.Printf("%s\t%d\n", s, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
