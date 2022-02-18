package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("esporte favorito: futebol")
	case "volei":
		fmt.Println("esporte favorito: volei")
	case "basquete":
		fmt.Println("esporte favorito: basquete")
	default:
		fmt.Println("sem esporte favorito")
	}
}
