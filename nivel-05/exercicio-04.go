package main

import "fmt"

func main() {
	x := struct {
		m map[string]int
		s []int
	}{
		map[string]int{
			"primeiro-item": 1,
			"segundo-item":  2,
			"terceiro-item": 3,
		},
		[]int{
			1, 2, 3, 4, 5,
		},
	}

	fmt.Println(x)
}
