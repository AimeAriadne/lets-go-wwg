package main

import (
	"fmt"
)

// Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

func main() {
	numero := 9

	switch {
	case numero == 0:
		fmt.Println("o número é 0")
	case numero%2 == 0:
		fmt.Println("o número é múltiplo de 2")
	case numero%3 == 0:
		fmt.Println("o número é múltiplo de 3")
	default:
		fmt.Println("O valor informado não atende as opções")
	}

}
