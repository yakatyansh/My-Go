package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch in Go")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Work hard")

	case 3:
		fmt.Println("Take a break")

	case 4:
		fmt.Println("Go to sleep")

	case 5:
		fmt.Println("Have a coffee")

	case 6:
		fmt.Println("Go for a walk")

	default:
		fmt.Println("What was that!")
	}

}
