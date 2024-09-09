package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// // kind of forEach loop
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v \n", index, day)
	// }

	// looping like while loop
	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			rougueValue++
			continue
		}

		if rougueValue == 5 {
			goto lco
		}

		if rougueValue == 7 {
			break
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

	// this is label for goto, continue and break
lco:
	fmt.Println("Welcome to LCO")
}
