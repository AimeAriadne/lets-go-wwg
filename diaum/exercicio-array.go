package main

import (
	"fmt"
)

// escreva um programa que contenha um array de strings.
// o valor de cada um dos elementos deve corresponder ao número do indice.
// print na tela o tipo do seu array e os valores de seus emelentos.

func main() {
	var numeros [6]string
	numeros[0] = "zero"
	numeros[1] = "um"
	numeros[2] = "dois"
	numeros[3] = "três"
	numeros[4] = "quatro"
	numeros[5] = "cinco"

	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros)
}
