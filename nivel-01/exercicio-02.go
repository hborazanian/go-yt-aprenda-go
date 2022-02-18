package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("x: %d, %T\n", x, x)
	fmt.Printf("y: %s, %T\n", y, y)
	fmt.Printf("z: %t, %T", z, z)
}
