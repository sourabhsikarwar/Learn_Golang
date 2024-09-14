package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "Hello Golang, I am working with files!"

	file, err := os.Create("./firstFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./firstFile.txt")
}

func readFile(fileName string) {
	// ioutil.ReadFile(fileName) --> This is deprecated, And this simply calls os.ReadFile, so using directly this
	dataByte, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data inside the file: ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
