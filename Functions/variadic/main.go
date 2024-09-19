package main

import "fmt"

func main() {
	// numbers
	sum := sumup(1, 10, 15, -6)

	//sum of list
	numbers := []int{10, 43, 12, 34, -13, -31}
	listSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(listSum)
}

//function with multiple arguements

func sumup(numbers ...int) int {

	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum

}
