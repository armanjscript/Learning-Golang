package main

import "fmt"

func main() {
	// var revenue int
	// var expenses int
	// var taxRate float64

	revenue := getUserInput("Revenue: ")

	expenses := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax Rate: ")

	// ebt := revenue - expenses
	// profit := float64(ebt) * (1 - taxRate/100)
	// ratio := float64(ebt) / profit

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.2f\n", ebt)
	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf("ratio: %.2f", ratio)

}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
