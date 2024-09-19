package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	firstname := getUserData("Please enter your first name: ")
	lastname := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate(MM/DD/YYYY): ")

	// var appUser user
	var appUser *user.User

	appUser, err := user.New(firstname, lastname, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	//... do sth awesome with that generated data!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin := user.NewAdmin("test@example.com", "test1234")
	admin.OutputUserDetails()
	admin.ClearUserName()
}

// func outputUserDetails(u *user) {
// 	// ...
// 	fmt.Println((*u).firstname, (*u).lastname, (*u).birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//Scanln for accepting enter key as input
	fmt.Scanln(&value)
	return value
}
