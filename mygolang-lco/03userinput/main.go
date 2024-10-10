package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to day 3"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a word : ")

	// comma ok syntax || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for writing, ", input)
	fmt.Printf("type of rating is %T", input)
}
