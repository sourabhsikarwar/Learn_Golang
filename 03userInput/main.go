package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input in Golang"
	fmt.Println(welcome)

	// Taking user input
	fmt.Println("Please rate out Golang input from 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this input is %T \n", input)
}
