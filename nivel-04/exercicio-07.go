// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {
	pessoas := [][]string{
		[]string{"Nome1", "Sobrenome1", "Hobby favorito1"},
		[]string{"Nome2", "Sobrenome2", "Hobby favorito2"},
		[]string{"Nome3", "Sobrenome3", "Hobby favorito3"},
	}

	for _, pessoa := range pessoas {
		fmt.Println(pessoa[0])
		for _, dado := range pessoa {
			fmt.Println("\t", dado)
		}
	}
}
