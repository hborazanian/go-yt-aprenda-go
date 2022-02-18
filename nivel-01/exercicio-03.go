package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("x: %d, %T\n", x, x)
	fmt.Printf("y: %s, %T\n", y, y)
	fmt.Printf("z: %t, %T\n", z, z)

	s := fmt.Sprintf("x: %d, y: %s, z: %t", x, y, z)

	fmt.Println(s)
}
