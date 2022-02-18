package main

import "fmt"

type myInt int

var x myInt

func main() {
	fmt.Printf("x: %v, %T\n", x, x)
	x = 42
	fmt.Printf("x: %v, %T", x, x)
}
