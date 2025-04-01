package main

import "fmt"

func add5NoReturn(a *int) {
	*a = *a + 5
}

func main () {
	x:=10
	add5NoReturn(&x);
	fmt.Println("After:", x)
}