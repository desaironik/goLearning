package main

import "fmt"
func main() {
	defer func() {
		if r := recover(); r !=nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before Panic")
	panic("Something went Wrong")
	fmt.Println("After Panic")
}