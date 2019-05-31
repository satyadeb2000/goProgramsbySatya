package main

import "fmt"

const (
	_ = iota
	x = 2020 + iota
	y = 2020 + iota
	z = 2020 + iota
	p = 2020 + iota
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(p)

}
