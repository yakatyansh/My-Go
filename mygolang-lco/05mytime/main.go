package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2023, time.January, 28, 21, 30, 00, 00, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
