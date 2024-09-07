package main

import (
	"fmt"
	"time"
)

// These are the default string must be provided for formatting the date or time in go
// Storing in constants for good practice
// NOTE: This is unusual but required
const DateFormat = "01-02-2006"
const DateWithTimeFormat = "01-02-2006 15:04:05"
const DateWithDayFormat = "01-02-2006 Monday"

func main() {
	fmt.Println("Welcome to time module in go!")

	timeFunc()
}

func timeFunc() {
	presentDate := time.Now()
	fmt.Println(presentDate)

	fmt.Println(presentDate.Format(DateFormat))

	createdDate := time.Date(2002, time.January, 18, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	// using Format with default format type strings stored as a constant
	fmt.Println(createdDate.Format(DateWithTimeFormat))
	fmt.Println(createdDate.Format(DateWithDayFormat))
}
