package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Files in Go")

	// Create a file
	f, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Write to a file
	l, err := f.WriteString("Hello World")
	if err != nil {
		panic(err)
	}
	fmt.Println(l, "bytes written successfully")

	// Read from a file
	f, err = os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bucket := make([]byte, 100)
	bytesread, err := f.Read(bucket)
	if err != nil {
		panic(err)
	}
	fmt.Println(bucket)
	fmt.Println(bytesread, "bytes read successfully")

}
