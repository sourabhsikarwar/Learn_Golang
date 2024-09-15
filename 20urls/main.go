package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://dummy-url:3000/sourabhsikarwar?name=sourabh&age=22"

func main() {
	fmt.Println("Welcome to Handling URLs module in golang")
	fmt.Println(myUrl)

	// Parsing the URL
	parsedURL, err := url.Parse(myUrl)
	checkNilErr(err)

	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)
	fmt.Println(parsedURL.Port())

	queryParams := parsedURL.Query()
	fmt.Println(queryParams)

	// returns an array of strings
	fmt.Println(queryParams["name"][0]) // sourabh
	fmt.Println(queryParams["age"])     // [22]

	for key, val := range queryParams {
		fmt.Printf("Key: %v, Value: %v\n", key, val[0])
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "name=sourabh&age=22",
	}

	fmt.Printf("Type of this URL is: %T\n", partsOfUrl) // type is url.URL

	finalUrl := partsOfUrl.String()
	fmt.Println(finalUrl) // type is converted to string
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
