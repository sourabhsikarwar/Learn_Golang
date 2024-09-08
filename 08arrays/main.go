package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	// Always need to provide the size for the array
	var carList [4]string
	carList[0] = "Merc"
	carList[1] = "BMW"
	carList[3] = "Mustang"

	// carList[2] will occupy a space and
	// the size will always remain which was initialed no matter the elements inside the array
	fmt.Println("My Car List: ", carList)

	var bikeList = [4]string{"Honda", "Yamaha", "Kawasaki"}

	fmt.Println("My Bike List: ", bikeList)
	fmt.Println("Size of Bike List: ", len(bikeList))
}
