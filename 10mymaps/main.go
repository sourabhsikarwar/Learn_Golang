package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang!")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for ", languages["JS"])

	// this will print zeroed value
	// as for string zeroed value is an empty string
	fmt.Println("Non initialized value: ", languages["RBI"])

	// Deleting values from maps

	// delete(languages, "RB")
	// fmt.Println("List of all languages: ", languages)

	// loops in golang
	// we can use _ in place of key or value to ignore it
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v \n", key, value)
	}
}
