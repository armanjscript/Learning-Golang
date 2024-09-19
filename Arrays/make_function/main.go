package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	//make string list with Make function - maximum length of items is 5 (capacity)
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//create map with make function
	courseRatings := make(floatMap, 3)

	courseRatings["react"] = 4.6
	courseRatings["node"] = 4.7
	courseRatings["angular"] = 4.4

	courseRatings.output()

	//For loop for arrays
	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

}
