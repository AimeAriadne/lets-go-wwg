package main

import (
	"fmt"
)

//Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade, maior de idade ou idosa.
//Leve em consideração os seguintes valores:
//Faixa
//Intervalos
//Menor de idade
//abaixo de 18
//Maior de idade
//entre 18 e 60
//Idoso
//acima de 60

func main() {
	var idade int
	fmt.Println("Qual é a sua idade?")
	fmt.Scan(&idade)
	if idade <= 18 {
		fmt.Println("Você é menor de idade")
	} else if idade >= 18 && idade <= 60 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é uma pessoa idosa")
	}
}
