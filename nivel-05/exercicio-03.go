// - Usando os structs veículo, caminhonete e sedan:
//     - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//     - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// - Demonstre estes valores.
// - Demonstre um único campo de cada um dos dois.
package main

import "fmt"

type Veiculo struct {
	Portas int
	Cor    string
}

type Caminhonete struct {
	Veiculo   Veiculo
	Tracao4x4 bool
}

type Sedan struct {
	Veiculo    Veiculo
	ModeloLuxo bool
}

func main() {
	caminhonete := Caminhonete{
		Veiculo: Veiculo{
			Portas: 2,
			Cor:    "preto",
		},
		Tracao4x4: true,
	}

	sedan := Sedan{
		Veiculo: Veiculo{
			Portas: 5,
			Cor:    "branco",
		},
		ModeloLuxo: false,
	}

	fmt.Println("Caminhonete:", caminhonete)
	fmt.Println("Sedan:", sedan)

	fmt.Println("Caminhonete Cor:", caminhonete.Veiculo.Cor)
	fmt.Println("Sedan Cor:", sedan.Veiculo.Cor)
}
