package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	basicSlice()
	deleteElement()
}

func basicSlice() {
	var fruitList = []string{"Apple", "Banana", "Watermelon"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// this will add orange to the list as it is a slice
	fruitList = append(fruitList, "Orange", "Grapes")
	fmt.Println(fruitList)

	// slicing the slice in golang

	// fruitList[1:] means from index 1 till the end
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// fruitList[:3] means from start till index 3
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	// different syntax for slices
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 745
	highScores[2] = 768
	highScores[3] = 456
	// highScores[4] = 900 // this will give an error as it is out of bound
	fmt.Println(highScores)

	// But with append we can add elements to the slice as it will reallocate the memory
	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	// sorting the slice in golang
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}

func deleteElement() {
	fmt.Println("Deleting elements from slices with an Index")

	var courses = []string{"golang", "reactjs", "javascript", "java", "python"}
	fmt.Println(courses)

	removeIndex := 2

	// removing an element from the slice based on the index
	courses = append(courses[:removeIndex], courses[removeIndex+1:]...)
	fmt.Println(courses)
}
