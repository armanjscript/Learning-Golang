package main

import (
	"fmt"
	"math"
)

// constant variable
const inflationRate = 2.5

func main() {

	//defining two variables together
	//var investmentAmount, years float64 = 1000, 10

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// expectedReturnRate := 5.5
	// var years float64 = 10

	//if we want to convert variable to any data types, we use shorter form :=
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// outputs information
	// fmt.Println(futureValue)
	// fmt.Printf("Future Value: %v \nFuture Value (adjusted fro inflation): %v", futureValue, futureRealValue)

	//Float value outputs with one decimal point
	// fmt.Printf("Future Value: %.1f \nFuture Value (adjusted fro inflation): %.1f", futureValue, futureRealValue)

	// fmt.Println("Future Value (adjusted fro inflation):", futureRealValue)

	//using multiline strings
	// 	fmt.Printf(`Future Value: %.1f,
	// Future Value (adjusted fro inflation): %.1f`, futureValue, futureRealValue)

	//Store value in string variable and print
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted fro inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
