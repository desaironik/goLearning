package main

import "fmt"

func add(a float32, b float32) float32 {
	return a + b
}

func main() {
	result := add(6.10, 7.90);
	fmt.Println(result)
}