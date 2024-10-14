package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("URLs in Go")

	// URL is a string

	myurl := "https://jsonplaceholder.typicode.com/posts"

	fmt.Println("URLS IN GO ")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Query Params: %T\n", qparams)

}
