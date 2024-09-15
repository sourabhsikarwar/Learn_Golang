package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://sourabhsikarwar.vercel.app"

func main() {
	fmt.Println("Welcome to web requests in golang")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close()

	fmt.Printf("Status code and Headers: %T\n %T\n", response.StatusCode, response.Header)

	data, err := io.ReadAll(response.Body)

	checkNilErr(err)

	content := string(data)

	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
