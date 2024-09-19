package main

import "fmt"

type transformFN func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers1 := transformNumbers(&numbers, transformerFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)

}

func transformNumbers(numbers *[]int, transform transformFN) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

//return functions as values
func getTransformerFunction(numbers *[]int) transformFN {

	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
