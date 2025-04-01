package main

import "fmt"

func main () {
	var stringArray = [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(len(stringArray))
	for index, stringChar := range stringArray {
		fmt.Println(index, stringChar)
	  }
}