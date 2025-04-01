package main

import "fmt"

func main () {
	var stringArray = [5]string{"A", "B", "C", "D", "E"}
	stringSlice := stringArray[0:5] 
	fmt.Println(len(stringArray))
	for index, stringChar := range stringArray {
		fmt.Println(index, stringChar)
	}


	stringSlice = append(stringSlice, "F")
	fmt.Println(len(stringSlice))
	for index, stringChar := range stringSlice {
		fmt.Println(index, stringChar)
	}  

}