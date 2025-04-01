package main

import "fmt"
func main() {
	arrVariable := []string{"ABCD", "Hello", "World"}
	for i,j:=range arrVariable {
		fmt.Println(i,j)
	}
}