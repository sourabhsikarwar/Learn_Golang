package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang")

	var num int = 20

	if num < 10 {
		fmt.Println("Number is less than 10")
	} else if num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is equal to 10")
	}

	// Can control the flow on the go
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Can initialize the value and then check with if else

	if age := 16; age >= 18 {
		fmt.Println("You can vote in India")
	} else {
		fmt.Println("You can't vote in India")
	}
}
