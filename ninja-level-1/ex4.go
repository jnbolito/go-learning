package main

import "fmt"

func main() {
	type jason int
	var x jason
	fmt.Printf("Type of x is %T and has value: %v\n", x, x)
	x = 42
	fmt.Println(x)
}
