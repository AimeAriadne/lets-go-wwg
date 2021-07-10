package main

import (
	"fmt"
)

//Utilizando a palavra reservada var declare uma variável numérica do tipo int.
//Atribua um valor a essa variável.
//Printe na tela o seu valor.
//Repita para 3 variáveis diferentes.

func main() {
	var num1 int = 100
	fmt.Println(num1)

	var num2 int = 200
	fmt.Printf("Meu tipo é %T e meu valor é: %v\n", num2, num2)

	var num3 int = 300
	fmt.Println(num3)

}
