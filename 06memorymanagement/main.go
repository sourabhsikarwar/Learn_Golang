package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Welcome to Memory Management in go!")

	cpuCount := runtime.NumCPU()
	fmt.Println("CPU Count: ", cpuCount)
}
