
package main

import (
       "fmt"
       "io"
       "os"
       "log"
       "crypto/sha256"
       "crypto/sha512"
       "hash"
)

func main() {
        var hasher hash.Hash

        if len(os.Args) > 1 {
	    if os.Args[1] == "384" {
	        hasher = sha512.New384()
	    } else if os.Args[1] == "256" {
	        hasher = sha256.New()
	    }
	}
	if hasher == nil {
            hasher = sha512.New()
        }
	fmt.Printf("%T\n", hasher)
	if _, err := io.Copy(hasher, os.Stdin); err != nil {
	   log.Fatal(err)
	}
        fmt.Printf("%x\n", hasher.Sum(nil))
}
