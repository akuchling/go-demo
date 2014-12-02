package main

// Better formatting of numbers

import (
	"fmt"
	"math"
)

var (
	MONTHS = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
)

type Money float64
type AmortizationSched []Money

func balance(L Money, n int, c float64, p int, extra Money) Money {
	nf := float64(n)
	pf := float64(p)
	// XXX wow, this is a lot of casts! Is there a cleaner way?
	bal := (float64(L)*(math.Pow(1+c, nf)-math.Pow(1+c, pf))/
		(math.Pow(1+c, nf)-1) - pf*float64(extra))
	// Round to 2 decimal places
	bal = math.Trunc(bal*100) / 100
	bal = math.Max(bal, 0.0)
	return Money(bal)
}

func time_to_payoff(amortization AmortizationSched) int {
	// Return the payment period in which the mortgage balance becomes 0.
	for i := 0; i < len(amortization); i++ {
		if amortization[i] <= 0 {
			return i
		}
	}
	return len(amortization)
}

func interest(i int, P Money, amortization AmortizationSched) Money {
	return P - (amortization[i] - amortization[i+1])
}

func total_interest(P Money, amortization AmortizationSched) Money {
	var total Money = 0.0
	for i := 0; i < len(amortization)-1; i++ {
		bal := amortization[i]
		if bal == 0 {
			break
		}
		total += interest(i, P, amortization)
	}
	return total
}

// XXX is there a Go built-in for this purpose?
// XXX how do I make this work on slices of arbitrary size?
func find(haystack [12]string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return i
		}
	}
	// XXX this should probably signal an error
	return len(haystack)
}

func main() {
	var L Money = 345600.0
	n := 12 * 30
	c := 4.875 / 12 / 100
	var extra Money = 120.0

	fmt.Println("Current loan")
	var payment Money = 2488.16
	current_month := 6

	// Compute amortization schedule
	amortization := make(AmortizationSched, n, n)
	for i := 0; i < n; i++ {
		amortization[i] = balance(L, n, c, i, extra)
	}

	fmt.Println("Months to payoff=", time_to_payoff(amortization))
	total_int := total_interest(payment, amortization)
	total_paid := L + total_int
	fmt.Println("Interest paid over lifetime=", total_int)
	fmt.Println("Total paid=", total_paid)

	fmt.Println("Month\tPrincipal\t%")
	var star string;
	month := "Sep"
	year := 2013

	for i := 0; i < len(amortization); i++ {
		if (i + 1 == current_month) {
			star = "*"
		} else {
			star = " "
		}

		fmt.Println(i+1, "\t", amortization[i], "\t",
			amortization[i] / L * 100, "\t", star, month, year)

		// Increment to next month, and increase the year if necessary.
		index := find(MONTHS, month)
		index = (index + 1) % len(MONTHS)
		month = MONTHS[index]
		if index == 0 {
			year += 1
		}
	}

}
