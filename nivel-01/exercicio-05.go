package main

import "fmt"

type myInt int

var x myInt
var y int

func main() {
	fmt.Printf("x: %v, %T\n", x, x)
	x = 42
	fmt.Printf("x: %v, %T\n", x, x)
	y = int(x)
	fmt.Printf("y: %v, %T", y, y)
}
