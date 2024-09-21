package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` // Name alias for json response
	Price    int      // Not mandotory to give alias
	Platform string   `json:"website"`
	Password string   `json:"-"`              // Hide the password from response of json with -
	Tags     []string `json:"tags,omitempty"` // omit the tags if it is empty
}

func main() {
	fmt.Println("Welcome to json in golang")
	encodeJson()
}

func encodeJson() {
	myCourse := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"VueJS Bootcamp", 399, "LearnCodeOnline.in", "abc123", nil},
	}

	finalJson, _ := json.MarshalIndent(myCourse, "", "\t") // 2nd params is prefix and 3rd is indentation

	fmt.Printf("%s\n", finalJson)
}
