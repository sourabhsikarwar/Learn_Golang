package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance, super or parent concept in golang

	sourabh := User{"Sourabh", "sourabhchenu@gmail.com", true, 22}
	fmt.Println(sourabh)
	fmt.Printf("Sourabh details are %+v\n", sourabh) // %+v is used to print the struct in a readable format

	fmt.Printf("Name is %v and Email is %v\n", sourabh.Name, sourabh.Email)
}
