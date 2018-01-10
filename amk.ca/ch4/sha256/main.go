
package main

import (
       "fmt"
       "crypto/sha256"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
        for i := range pc {
                pc[i] = pc[i/2] + byte(i&1)
        }
}

// PopCount returns the population count (number of set bits) of x.
func popCount(x uint8) int {
     return int(pc[byte(x)])
}

// Count differing bits in two digests
func countBits(c1 [32]byte, c2[32]byte) int {
     result := 0
     for i := 0; i < 32; i++ {
     	 result += popCount(c1[i] ^ c2[i])
     }
     return result
}

//!-

func main() {
        c1 := sha256.Sum256([]byte("x"))
        c2 := sha256.Sum256([]byte("X"))

        fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1,
	   countBits(c1, c2),

	)
        // Output:
        // 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
        // 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
        // false
        // [32]uint8
}
