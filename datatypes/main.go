package main

import "fmt"
func main() {
	var num int = 42
	var ptr *int = &num
	fmt.Printf("Memory address of ptrL %d\n", ptr)
	fmt.Printf("Value of ptr %d\n", *ptr)
}