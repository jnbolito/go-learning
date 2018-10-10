package main

import "fmt"

const yearBase = 2015
const (
	year1 = yearBase + iota
	year2 = yearBase + iota
	year3 = yearBase + iota
	year4 = yearBase + iota
)

func main() {
	fmt.Println(year1, year2, year3, year4)
}
