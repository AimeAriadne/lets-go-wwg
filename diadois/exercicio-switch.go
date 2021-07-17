package main

import (
	"fmt"
)

// declarar uma varíavel chamada hora e atribuir um valor int a ela.
// usando switch, elencar cases e printar na tela se a hora corresponde ao período da manhã, tarde, noite ou madrugada.

func main() {
	hora := 0

	switch {
	case hora >= 6 && hora <= 11:
		fmt.Printf("São %d e horas está de manhã", hora)
	case hora >= 12 && hora <= 17:
		fmt.Printf("São %d horas e está de tarde", hora)
	case hora >= 18 && hora <= 23:
		fmt.Printf("São %d horas e está de noite", hora)
	case hora >= 0 && hora <= 5:
		fmt.Printf("São %d horas e está de madrugada", hora)
	default:
		fmt.Println("O valor informado não é uma hora válida")

	}

}
