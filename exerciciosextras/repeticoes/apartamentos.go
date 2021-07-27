package main

import (
	"fmt"
)

//Considerando os tópicos que já aprendemos até agora: slices, structs ,condicionais e laços de repetição, crie um programa que traga as informações sobre apartamentos de um prédio. Passos:
//Crie uma estrutura que representa um apartamento, com campos para representar seu número, o nome da sua proprietária e se tem vaga de garagem
//Reúna as estruturas em uma slice que representa um conjunto de apartamentos
//Printe as informações de cada unidade, separando por linha, usando for range

type Apartment struct {
	number int
	owner  string
	garage bool
}

func main() {
	apart1 := Apartment{
		number: 666,
		owner:  "Cruella",
		garage: true,
	}
	apart2 := Apartment{
		number: 777,
		owner:  "Bay",
		garage: false,
	}
	apart3 := Apartment{
		number: 222,
		owner:  "Caetana",
		garage: true,
	}
	apart4 := Apartment{
		number: 111,
		owner:  "Joy",
		garage: false,
	}

	apartments := []Apartment{apart1, apart2, apart3, apart4}

	for _, apartment := range apartments {
		parkingSlot := "it does not have a parking slot"
		if apartment.garage {
			parkingSlot = "it has a parking slot"
		}
		fmt.Printf("Apartment %d belongs to %s and %s\n", apartment.number, apartment.owner, parkingSlot)
	}
}
