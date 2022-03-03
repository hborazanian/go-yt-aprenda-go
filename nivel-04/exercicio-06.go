package main

import "fmt"

func main() {

	estados := make([]string, 26)

	fmt.Println(len(estados), cap(estados))

	estados = []string{
		"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás",
		"Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais",
		"Pará", "Paraíba", "Paraná", "Pernambuco",
		"Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins",
	}

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
