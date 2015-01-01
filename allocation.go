// Takes standard input containing an investment portfolio and set of
// prices, and calculates the total value and the percentage
// allocations.
//
//
// Input file:
//    <fund symbol> <units>
//    <fund symbol 2> <units>
//    ...
//    <blank line>
//    <fund symbol> <type> <price>
//    <fund symbol 2> <type 2> <price 2>
//    ...
//
// The funds need not be in the same order in the two sections of file.
// The script will report errors if a fund's price isn't given or is
// duplicated.  The same symbol can occur multiple times in the first
// section.

package main

import (
	"fmt"
)

type Stock struct {
	price float64
	fund_type string
}
type StockDatabase map[string]Stock

type Position struct {
	symbol string
	units float64
}


func output_report (stocks StockDatabase, portfolio []Position) {
	var total float64
	total_by_symbol := make(map[string]float64)
	total_by_type := make(map[string]float64)

	for i := 0; i < len(portfolio); i++ {
		symbol := portfolio[i].symbol
		record, ok := stocks[symbol]
		if !ok {
			// Record not found
			// XXX how to print to standard error?
			_, err := fmt.Println("No price found for symbol",
	 			              symbol)
			if (err != nil) {
				// XXX how to handle?
			}
			continue
		}
		amt := portfolio[i].units * record.price
		total_by_symbol[portfolio[i].symbol] += amt
		total_by_type[record.fund_type] += amt;
		total += amt
	}


	fmt.Println("Totals by fund")
	// XXX how to do this in sorted order?
	for symbol, fund_total := range total_by_symbol {
		fmt.Printf("%-15s\t$%10.2f\t%.1f%%\n", symbol, fund_total,
			fund_total / total * 100.0)
	}
	fmt.Println("")

	fmt.Println("Fund types")
	// XXX how to do this in sorted order?
	for fund_type, type_total := range total_by_type {
		fmt.Printf("%-15s\t$%10.2f\t%.1f%%\n", fund_type, type_total,
			type_total / total * 100.0)

	}
	fmt.Println("")

}

func main() {
	var stocks StockDatabase
	var portfolio []Position

	stocks = make(StockDatabase)

	// Test code
	stocks["STB1"] = Stock{price: 10.00, fund_type: "bond"}
	stocks["STB2"] = Stock{price: 15.00, fund_type: "bond"}
	stocks["STL1"] = Stock{price: 20.00, fund_type: "large-cap"}

	portfolio = append(portfolio, Position{symbol: "STB1", units: 10})
	portfolio = append(portfolio, Position{symbol: "STB1", units: 25})
	portfolio = append(portfolio, Position{symbol: "STB2", units: 35})
	portfolio = append(portfolio, Position{symbol: "STL1", units: 40})

	output_report(stocks, portfolio)
}
