package main

import (
	"fmt"
)

//Construa uma função que receba uma lista de números inteiros, modifique essa lista dobrando os números ímpares e divida por dois os pares, e retorne a lista modificada e a soma de todos os elementos da lista.
var numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var total int = 0

func main() {
	for i, v := range numeros {
		switch {
		case (v % 2) != 0:
			numeros[i] *= 2

		case (v % 2) == 0:
			numeros[i] /= 2
		}
		total += numeros[i]
	}
	fmt.Printf("%v\n%d", numeros, total)
}
