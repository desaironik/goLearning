package main

import "fmt"
func main() {
	count := 6

	switch {
	case count % 2 == 0:
		fmt.Println("Divisible By 2")
		
	case count%3 == 0:
		fmt.Println("Divisible By 3")
	default:
		fmt.Println("Not Divisible by 2 or 3")
	}	
}