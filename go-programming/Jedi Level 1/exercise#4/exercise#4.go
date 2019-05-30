package main

import (
	"fmt"
)

type satya int

var x satya

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
