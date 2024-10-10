package main

import "fmt"

func main() {

	fmt.Println("Pointers in Go")
	// var ptr *int
	// fmt.Println("Value of the pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the pointer is ", ptr)
	fmt.Println("Value of the pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of the pointer is ", *ptr)     // 46
	fmt.Println("Value of the pointer is ", myNumber) // 46

	// Pointer to a pointer
	var ptr1 = &ptr
	fmt.Println("Value of the pointer is ", ptr1)   // address of ptr
	fmt.Println("Value of the pointer is ", *ptr1)  // address of myNumber
	fmt.Println("Value of the pointer is ", **ptr1) // 46

}
