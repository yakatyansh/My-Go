package main

import "fmt"

func main() {

	logincount := 23
	var result string
	fmt.Println("If Else in Go")

	if logincount < 10 {
		result = "Regular User"

	} else if logincount > 10 && logincount < 100 {
		result = "Special User"
	}
	fmt.Println("User Status: ", result)

	if logincount > 100 {
		result = "Super User"
	}

	fmt.Println("User Status: ", result)

}
