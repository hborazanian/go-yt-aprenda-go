package main

import "fmt"

func main() {
	// 08
	pessoas := map[string][]string{
		"pessoa-01": []string{
			"hobbie01",
			"hobbie02",
			"hobbie03",
			"hobbie04",
			"hobbie05",
			"hobbie06",
		},
		"pessoa-02": []string{
			"hobbie01",
			"hobbie02",
			"hobbie03",
		},
	}

	for pessoa, hobbies := range pessoas {
		fmt.Println(pessoa)
		for index, hobbie := range hobbies {
			fmt.Println("\t", index, hobbie)
		}
	}

	// 09
	pessoas["pessoa-03"] = []string{
		"hobbie01",
		"hobbie02",
	}

	for pessoa, hobbies := range pessoas {
		fmt.Println(pessoa)
		for index, hobbie := range hobbies {
			fmt.Println("\t", index, hobbie)
		}
	}

	// 10
	delete(pessoas, "pessoa-02")

	for pessoa, hobbies := range pessoas {
		fmt.Println(pessoa)
		for index, hobbie := range hobbies {
			fmt.Println("\t", index, hobbie)
		}
	}
}
