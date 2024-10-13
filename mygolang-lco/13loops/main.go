package main

import "fmt"

func main() {

	fmt.Println("Loops in Go")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	routineloop := 1
	for routineloop < 5 {

		if routineloop == 3 {
			goto lco
		}

		if routineloop == 2 {
			fmt.Println("Value is 2, so skipping")
			routineloop++
			continue
		} else if routineloop == 4 {
			fmt.Println("Value is 4, so breaking")
			break
		}
		fmt.Println("Value is ", routineloop)
		routineloop++
	}

lco:
	fmt.Println("Jumping to LCO")

}
