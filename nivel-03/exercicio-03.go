package main

import (
	"fmt"
	"time"
)

func main() {
	yearBirth := 1989
	yearCurrent := time.Now().Year()
	for yearBirth <= yearCurrent {
		fmt.Println(yearBirth)
		yearBirth++
	}
}
