package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slices in Go")
	var fruitlist = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "Mango", "Peach")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	highscores = append(highscores, 555, 666, 777)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)

	// remove a value from slice based on index

	var index int = 2
	highscores = append(highscores[:index], highscores[index+1:]...)
	fmt.Println(highscores)

}
