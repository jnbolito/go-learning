package main

import "fmt"

func main() {
	const a = 2
	printInt(a)
	printInt(a << 1)
}

func printInt(a int) {
	fmt.Printf("%v %b %#x\n", a, a, a)
}
