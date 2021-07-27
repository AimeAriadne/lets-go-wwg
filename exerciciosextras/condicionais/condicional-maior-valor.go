package main

import (
	"fmt"
)

// Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.

func main() {
	idade1 := 27
	idade2 := 29
	idade3 := 25

	if idade1 > idade2 && idade1 > idade3 {
		fmt.Printf("A maior idade é %d", idade1)
	} else if idade2 > idade1 && idade2 > idade3 {
		fmt.Printf("A maior idade é %d", idade2)
	} else if idade3 > idade1 && idade3 > idade2 {
		fmt.Printf("A maior idade é %d", idade3)

	}
}
