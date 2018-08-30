package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World, this is GO! Coming to you live from SAN TESTINGO!")
	result := addTwo(2, 5)
	fmt.Println(result)
}

func addTwo(x, y int) (result int) {
	result = x + y
	return result
}
