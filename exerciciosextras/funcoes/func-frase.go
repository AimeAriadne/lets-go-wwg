package main

import (
	"fmt"
	"strings"
)

//Construa uma função que receba uma palavra ou frase e uma letra, e retorne o número de ocorrências da letra informada.

func main() {
	frase := "fora bolsonaro"
	letra := "o"

	var resultado = strings.Count(frase, letra)
	fmt.Printf("A letra '%s' foi encontrada %d vezes na frase '%s'", letra, resultado, frase)
}
