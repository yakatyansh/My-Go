package main

import "fmt"

func main() {

	fmt.Println("Defer in Go")
	defer fmt.Println("End")
	fmt.Println("Start")
	mydefer()

}

func mydefer() {
	defer fmt.Println("End of mydefer")
	fmt.Println("Start of mydefer")
}
