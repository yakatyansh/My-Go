package main

import "fmt"

func main() {
	var username string = "Yash"
	fmt.Println(username)
	fmt.Printf("The Variable is of Type : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("The Variable is of Type : %T \n", isLoggedin)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("The Variable is of Type : %T \n", smallVal)

	var smallFloat float64 = 23.877878990099
	fmt.Println(smallFloat)
	fmt.Printf("The Variable is of Type : %T \n", smallFloat)

	//default values
	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Printf("The Variable is of Type : %T \n", anotherVar)

	// implicit type declaration
	var website = "upsc.gov.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
