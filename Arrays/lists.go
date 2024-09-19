package main

import "fmt"

func main() {
	//Dynamic lists or arrays
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])

	// updated_prices := append(prices, 3.86)
	// fmt.Println(updated_prices, " ", prices)
	prices = append(prices, 3.86)
	fmt.Println(prices)
}

// func main() {
// 	//defining array
// 	prices := [5]float64{10.99, 9.99, 45.99, 20.0, 34.99}
// 	fmt.Println(prices)

// 	//fixed array without initialization
// 	var productNames [4]string
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2:])
// 	fmt.Println(prices[:2])
// 	fmt.Println(cap(prices[:2]), " ", len(prices[:2]))

// }
