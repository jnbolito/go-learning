package main

import "fmt"

func main() {
	a := 2 == 3
	b := 2 < 3
	c := 2 > 3
	d := 2 >= 3
	e := 23 <= 2
	f := 2 != 3

	fmt.Println(a, b, c, d, e, f)
}
