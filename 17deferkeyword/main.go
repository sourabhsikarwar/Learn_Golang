package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer keyword in golang")

	defer fmt.Println("World")
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Start...")

	myDefer()
}

// Works in LIFO order i.e. last in first out
// defer keyword makes the line of code to be executed after the function returns

// World Hello One Two Three 0 1 2 3 4
// the order of insertion of defer keyword is LIFO

// output:
// Start...
// 4
// 3
// 2
// 1
// 0
// Three
// Two
// One
// Hello
// World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
