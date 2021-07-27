package main

import (
	"fmt"
	"os"
)

//Faça um programa que solicite à usuária que informe um número até que o número informado seja par.

func main() {
	var numero = 1

	for numero%2 != 0 {
		fmt.Println("Insira um número:")
		fmt.Fscanf(os.Stdin, "%d", &numero)
	}
	fmt.Println("Agora sim podemos dividir igualmente entre mim e você!")

}
