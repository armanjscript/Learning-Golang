package main

import "fmt"

func main() {
	//defining maps
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon"])

	//assign value to new key
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	//delete a key and its value
	delete(websites, "Google")
	fmt.Println(websites)
}
