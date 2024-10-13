package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greet()
	fmt.Println(add(2, 3))
	fmt.Println(proAdder(2, 3, 4, 5))
}

func greet() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// Output
// Functions in Go
// Hello, World!
// 5
// Program exited.
