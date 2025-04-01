package main

import "fmt"
func main() {
	age := 78
	if age > 25 && age < 62 {
		fmt.Println("Age Greater than 25")
	} else if age>= 62 {
		fmt.Println("Age Greater than 62")
	} else {
		fmt.Println("Age less than 25")
	}
}