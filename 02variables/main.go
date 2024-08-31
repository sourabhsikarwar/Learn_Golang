package main

import "fmt"

// Only allowed method for variable decl. outside block or method
var globalVar string = "Hello World"

// Not allowed (Only allowed inside methods)
// globalVar := "Hello World"

// -----------

// with Capital letter at first position in the variable name,
// this variable is now public and all the files can access it
const DummyToken = "DummyToken"

func main() {
	var username string = "Sourabh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var num int = 65
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	var floatNum float32 = 65.56758484848
	fmt.Println(floatNum)
	fmt.Printf("Variable is of type: %T \n", floatNum)

	var largeFloatNum float64 = 65.567584848487585588
	fmt.Println(largeFloatNum)
	fmt.Printf("Variable is of type: %T \n", largeFloatNum)

	// Default values
	var num1 int
	fmt.Println(num1)
	fmt.Printf("Variable is of type: %T \n", num1)

	var str1 string
	fmt.Println(str1)
	fmt.Printf("Variable is of type: %T \n", str1)

	// implicit declaration
	var str2 = "Hello"
	fmt.Println(str2)
	fmt.Printf("Variable is of type: %T \n", str2)

	// no var keyword
	str3 := "Hello with No Var Keyword"
	fmt.Println(str3)
	fmt.Printf("Variable is of type: %T \n", str3)

	fmt.Println(globalVar)

	fmt.Println(DummyToken)
	fmt.Printf("Variable is of type: %T \n", DummyToken)

	// Multiple variable declaration

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
}
