package main

import "fmt"

func min(vals ...int) (result int) {
	result = vals[0]

	for _, v := range vals {
		if v < result {
			result = v
		}
	}
	return
}

func max(vals ...int) (result int) {
	result = vals[0]

	for _, v := range vals {
		if v > result {
			result = v
		}
	}
	return
}

func main() {
	figures := []int{3, 4, 20, -17, 42}

	fmt.Println("List=", figures)

	fmt.Println("max =", max(figures...))
	fmt.Println("min =", min(figures...))
}
