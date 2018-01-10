// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate_onepass(s, 3)
	fmt.Println(s) // "[3 4 5 0 1 2]"
	//!-slice

}

func rotate(s []int, n int) {
     reverse_slice(s[:n])
     reverse_slice(s[n:])
     reverse_slice(s)
}

func rotate_onepass(s []int, n int) {
     var result []int

     size := len(s)
     for i := 0; i < size; i++ {
     	 j := (i - n)
	 if j < 0 {
	    j += size
	 }
	 result = append(result, s[j])
     }
     for i := 0; i < size; i++ {
     	 s[i] = result[i]
     }
}


// reverse reverses a slice of ints in place.
func reverse_slice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
