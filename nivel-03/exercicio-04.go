package main

import (
	"fmt"
	"time"
)

func main() {
	yearBirth := 1989
	yearCurrent := time.Now().Year()

	for {
		if yearBirth > yearCurrent {
			break
		}

		fmt.Println(yearBirth)
		yearBirth++
	}
}
