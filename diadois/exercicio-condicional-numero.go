package main

import (
	"fmt"
)

//Declara uma variável;
//Atribui o valor de um número a ela;
//Informa se o número é positivo ou negativo.

func main() {
	num := -21
	if num >= 0 {
		fmt.Printf("O número %d é positivo", num)
	} else {
		fmt.Printf("O número %d é negativo", num)
	}

}
