package main

// Better formatting of numbers
// Add a Money type

import (
	"fmt"
	"math"
)

var (
	MONTHS = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
)

type AmortizationSched []float64

func balance(L float64, n int, c float64, p int, extra float64) float64 {
	nf := float64(n)
	pf := float64(p)
	bal := (L*(math.Pow(1+c, nf)-math.Pow(1+c, pf))/
		(math.Pow(1+c, nf)-1) - pf*extra)
	// Round to 2 decimal places
	bal = math.Trunc(bal*100) / 100
	bal = math.Max(bal, 0.0)
	return bal
}

func time_to_payoff(amortization []float64) int {
	// Return the payment period in which the mortgage balance becomes 0.
	for i := 0; i < len(amortization); i++ {
		if amortization[i] <= 0 {
			return i
		}
	}
	return len(amortization)
}

func interest(i int, P float64, amortization AmortizationSched) float64 {
	return P - (amortization[i] - amortization[i+1])
}

func total_interest(P float64, amortization AmortizationSched) float64 {
	total := 0.0
	for i := 0; i < len(amortization)-1; i++ {
		bal := amortization[i]
		if bal == 0 {
			break
		}
		total += interest(i, P, amortization)
	}
	return total
}

func main() {
	L := 345600.0
	n := 12 * 30
	c := 4.875 / 12 / 100
	extra := 120.0

	fmt.Println("Current loan")
	payment := 2488.16
	//current_month := 6

	// Compute amortization schedule
	amortization := make(AmortizationSched, n, n)
	for i := 0; i < n; i++ {
		amortization[i] = balance(L, n, c, i, extra)
		//fmt.Println(i+1, amortization[i])
	}

	fmt.Println("Months to payoff=", time_to_payoff(amortization))
	total_int := total_interest(payment, amortization)
	total_paid := L + total_int
	fmt.Println("Interest paid over lifetime=", total_int)
	fmt.Println("Total paid=", total_paid)

}
