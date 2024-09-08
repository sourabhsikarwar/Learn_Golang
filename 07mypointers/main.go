package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers module")

	// initializing pointers in go
	// this will initialize the pointer to nil
	var ptr *int
	fmt.Println("Value of pointer is: ", ptr) // this will give <nil>

	// using new keyword
	myNumber := 20

	var numPtr = &myNumber

	fmt.Println("Memory Address of myNumber: ", numPtr)
	fmt.Println("Actual value of numPtr: ", *numPtr)

	// math operation with pointers
	*numPtr = *numPtr + 5
	fmt.Println("Value after operation: ", *numPtr)

	// With new keyword
	// this will initialize a pointer of type int with a valid memory address
	var ptr2 = new(int)
	fmt.Println("Value of pointer is: ", ptr2)
	fmt.Println("Value of pointer is: ", *ptr2) // this will give a value of 0

	*ptr2 = 10
	fmt.Println("Value of pointer is: ", *ptr2)
}
