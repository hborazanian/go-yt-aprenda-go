// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

package main

import "fmt"

type Pessoas []Pessoa

type Pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	// 01
	pessoas := Pessoas{}

	pessoa1 := Pessoa{
		Nome:      "Nome Pessoa 01",
		Sobrenome: "Sobrenome Pessoa 01",
		Sabores: []string{
			"sabor-01",
			"sabor-02",
			"sabor-03",
			"sabor-04",
		},
	}
	pessoas = append(pessoas, pessoa1)

	pessoa2 := Pessoa{
		Nome:      "Nome Pessoa 02",
		Sobrenome: "Sobrenome Pessoa 02",
		Sabores: []string{
			"sabor-01",
			"sabor-02",
		},
	}

	pessoas = append(pessoas, pessoa2)

	for _, pessoa := range pessoas {
		fmt.Println("Pessoa: ", pessoa.Nome, pessoa.Sobrenome)
		fmt.Println("\t Sabores:")
		for _, sabor := range pessoa.Sabores {
			fmt.Println("\t", sabor)
		}

		fmt.Println("")
	}

	// 02

	mapPessoas := map[string]Pessoa{}

	for _, pessoa := range pessoas {
		mapPessoas[pessoa.Sobrenome] = pessoa
	}

	for _, pessoa := range mapPessoas {
		fmt.Println("Pessoa: ", pessoa.Nome, pessoa.Sobrenome)
		fmt.Println("\t Sabores:")
		for _, sabor := range pessoa.Sabores {
			fmt.Println("\t", sabor)
		}

		fmt.Println("")
	}
}
