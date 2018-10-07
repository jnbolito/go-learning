package main

import "fmt"

type jason int
var y int

func main() {
	var x jason
	fmt.Printf("Type of x is %T and has value: %v\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("y is of type %T and has value: %v", y, y)
}
