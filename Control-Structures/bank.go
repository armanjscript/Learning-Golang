package main

import (
	"fmt"

	"example.com/bank/fileops"
	// Third Party Package
	// "github.com/grandper/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	// r := randomdata.FromSeed(1234)
	// fmt.Println(r.PhoneNumber())

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		//panic function for stopping execution
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid withdrawl amount. withdrawl more than account balance not allowed")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}

}
