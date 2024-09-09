package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	// this is deprecated
	// rand.Seed(time.Now().UnixNano())

	rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You got 1, You can move 1 spot forward or open")
	case 2:
		fmt.Println("You got 2, You can move 2 spot forward")
	case 3:
		fmt.Println("You got 3, You can move 3 spot forward")
		fallthrough // this will execute the next case
	case 4:
		fmt.Println("You got 4, You can move 4 spot forward")
	case 5:
		fmt.Println("You got 5, You can move 5 spot forward")
	case 6:
		fmt.Println("You got 6, You can move 6 spot forward or open with one more roll")
	default:
		fmt.Println("WTF!!!")
	}
}
