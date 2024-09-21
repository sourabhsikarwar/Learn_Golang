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
	// encodeJson()
	decodeJson()
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

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	isValidJson := json.Valid(jsonDataFromWeb)

	if isValidJson {
		fmt.Println("Json is valid")
		var myCourse course

		// decoding the json into a struct
		checkValid := json.Unmarshal(jsonDataFromWeb, &myCourse) // 2nd param is a pointer

		if checkValid == nil {
			fmt.Printf("%#v\n", myCourse) // %#v is used to print the struct in a readable format
		}
	} else {
		fmt.Println("Json is not valid")
	}

	var myOnlineClass map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineClass)
	fmt.Printf("%#v\n", myOnlineClass)

	for k, v := range myOnlineClass {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
