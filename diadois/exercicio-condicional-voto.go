package main

import (
	"fmt"
	"time"
)

//Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
//Declare uma variável e atribua a ela o ano atual.
//Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.

func main() {
	anoDeNascimento := 1994
	anoAtual := time.Now().Year()

	idade := anoAtual - anoDeNascimento

	if idade >= 16 {
		fmt.Printf("Essa pessoa tem %d anos e ela já pode votar.", idade)
		return
	}
	fmt.Printf("Essa pessoa tem %d anos e ela não pode votar.", idade)

}
