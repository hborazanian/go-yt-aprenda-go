package main

import "fmt"

func main() {
	x := 11

	if x == 10 {
		fmt.Println("x == 10")
	} else if x > 10 {
		fmt.Println("x > 10")
	} else {
		fmt.Println("x != 10 && !(x > 10)")
	}

	switch {
	case x == 10:
		fmt.Println("x == 10")
	case x > 10:
		fmt.Println("x > 10")
	default:
		fmt.Println("x != 10 && !(x > 10)")
	}
}
