package main

import "fmt"

func main() {
	start := 10
	stop := 100
	div := 4

	for i := start; i <= stop; i++ {
		fmt.Printf("n: %d - sobra: %d\n", i, (i % div))
	}
}
