package main

import "fmt"

func main() {
	x := 42
	y := 84
	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x != y)
	e := (x > y)
	f := (x < y)

	fmt.Println(a, "\t", b, "\t", c, "\t", d, "\t", e, "\t", f)

}
