package main

import (
	"fmt"
)

const (
	_ = 2022 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Print(a, b, c, d)
}
