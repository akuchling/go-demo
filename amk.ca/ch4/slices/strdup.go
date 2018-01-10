// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	s := []string{"aardvark", "bear", "bear", "cat", "cat", "cat", "dog"}
	remove_dups(s)
	fmt.Println(s)

}

func remove_dups(s []string) {
     size := len(s)
     var i, j int;
     for i, j = 0, 0; i < size; i++ {
         for i < size - 1 && s[i] == s[i+1] {
	     i++
	 }
	 fmt.Println("Copying", i, "to", j)
     	 s[j] = s[i]
	 j++
     }
     for i = j; i < size; i++ {
     	 s[i] = ""
     }
}
