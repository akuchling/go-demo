package main

import "fmt"

// Returns 42, but has no return statement

func panicfunc() (result int) {
     type bailout struct {value int}

     defer func() {
        p := recover()
	if p != nil {
	   result = p.(bailout).value
	}
     }()
     panic(bailout{42})
}

func main() {
     fmt.Println("Should return 42:", panicfunc())

}
