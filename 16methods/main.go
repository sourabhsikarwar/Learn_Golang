package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to methods in golang")

	// no inheritance, super or parent concept in golang

	sourabh := User{"Sourabh", "sourabhchenu@gmail.com", true, 22}
	fmt.Println(sourabh)
	fmt.Printf("Sourabh details are %+v\n", sourabh) // %+v is used to print the struct in a readable format

	fmt.Printf("Name is %v and Email is %v\n", sourabh.Name, sourabh.Email)

	// methods
	sourabh.getStatus()
	sourabh.changeMail("sourabhchenu@go.dev")

	// mail is not changed here as the user object passed is a copy
	// of the original user object
	// Thus pointers are used and useful
	fmt.Printf("Name is %v and Email is %v\n", sourabh.Name, sourabh.Email)

}

func (u User) getStatus() {
	fmt.Printf("User %v status is %v\n", u.Name, u.Status)
}

// Here u is a copy of the original user object
func (u User) changeMail(newMail string) {
	u.Email = newMail
	fmt.Printf("User %v email is %v\n", u.Name, u.Email)
}
