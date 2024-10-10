package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in Go")
	languages := make(map[string]string)
	languages["GO"] = "Go"
	languages["RU"] = "Ruby"
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	fmt.Println(languages)
	fmt.Println("GO stands for ", languages["GO"])

	delete(languages, "JS")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
