package main

import (
	"fmt"
)

func main() {
	x:=1980
	for{
		fmt.Println(x)
		x++
		if x>2019{
			break
		}
	}
}
