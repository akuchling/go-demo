// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var x IntSet
	fmt.Println(x.Len())

	x.Add(1)
	fmt.Println(x.Len())
	x.Add(2)
	x.Add(3)
	fmt.Println(x.Len())

	var y *IntSet
	fmt.Println(y.Len())

	// Output:
	// 0
	// 1
	// 3
	// 0
}

func TestRemoveClear(t *testing.T) {
	var x IntSet
	for i := 0; i < 277; i++ {
		x.Add(i)
	}
	fmt.Println(x.Len())

	x.Remove(1)
	fmt.Println(x.Len())
	x.Clear()
	fmt.Println(x.Len())
	fmt.Println(x)

	// Output:
	// 277
	// 276
	// 0
	// {[]}

}

func TestCopy(t *testing.T) {
	var x IntSet
	for i := 0; i < 277; i++ {
		x.Add(i)
	}
	fmt.Println(x.Len())

	y := x.Copy()
	x.Remove(1)
	fmt.Println(x.Len())
	fmt.Println(y.Len())

	y.Clear()
	fmt.Println(x.Len())
	fmt.Println(y.Len())

	// Output:
	// 277
	// 276
	// 277
	// 276
	// 0

}

func TestAddAll(t *testing.T) {
	var x IntSet

	x.AddAll(17, 117, 249)
	fmt.Println(x.Len())
	fmt.Println(x.String())

	// Output:
	// 3
	// {17 117 249}
}

func TestElems(t *testing.T) {
	var x IntSet

	x.AddAll(17, 432, 117, 249)
	fmt.Println(x.Elems())

	// Output:
	// [17 117 249 432]
}
