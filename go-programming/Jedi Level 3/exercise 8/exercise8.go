package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	hr, min, sec := timeStamp.Clock()

	switch {
	case hr <= 11:
		fmt.Println("Its Datime\t", hr, min, sec)
	default:
		fmt.Println("Goodnight", hr, min, sec)

	}

}
