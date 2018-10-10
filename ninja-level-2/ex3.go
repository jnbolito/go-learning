package main

import "fmt"

func main() {
	const a = 3
	const b int = 4

	fmt.Printf("%#v %T", a, a)
	fmt.Printf("%d %T", b, b)

}
