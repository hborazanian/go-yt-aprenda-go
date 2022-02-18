package main

import "fmt"

func main() {
	asciiA := 65
	asciiZ := 90
	repeat := 3

	for i := asciiA; i <= asciiZ; i++ {
		fmt.Println(i)
		for x := 0; x < repeat; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
