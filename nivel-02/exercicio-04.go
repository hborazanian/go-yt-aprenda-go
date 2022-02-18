package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("d: %d, b: %b, e: %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("d: %d, b: %b, e: %#x", y, y, y)
}
