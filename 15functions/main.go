package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	greeting()

	fmt.Println(multiplier(2, 8))

	// Spreading the params
	result := proMultiplier(2, 4, 8, 16)

	fmt.Println("Result is: ", result)

	value, message := proAdder(2, 4, 8, 16)

	fmt.Printf("Value is : %v and Message: %v \n", value, message)
}

// multiple return values
func proAdder(numbers ...int) (int, string) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total, "All the values added together!"
}

// spreading the args
func proMultiplier(values ...int) int {
	total := 1
	for _, val := range values {
		total *= val
	}
	return total
}

func multiplier(valOne int, valTwo int) int {
	return valOne * valTwo
}

func greeting() {
	fmt.Println("Hello there!")
}
