package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	var myArray [5]int
	fmt.Println(myArray)

	var fruitlist [10]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Orange"
	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

}
