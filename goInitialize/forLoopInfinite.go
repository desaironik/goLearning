package main

import "fmt"
func main() {
	var i int = 1
	for {
		fmt.Println(i)
		i++
		if (i >10) {
			break;
		}
	}
}