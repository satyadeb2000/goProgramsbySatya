package main

import "fmt"

const (
	a float64 = 43.2
	b         = 43
)

func main() {
	fmt.Printf("%v\t%v\n", a, b)
	fmt.Printf("%T\t%T\n", a, b)
}
