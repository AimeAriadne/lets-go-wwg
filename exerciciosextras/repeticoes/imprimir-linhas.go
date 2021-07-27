package main

import (
	"fmt"
)

//Faça um programa que imprima 10 linhas, onde cada linha obedece a seguinte regra: na linha 1 deve ser impresso um valor (1), na linha 2 devem ser impressos dois valores (1 e 2), na linha três devem ser impressos três valores (1, 2, e 3), … na linha 10 devem ser impressos dez valores (1, 2, 3, 4, 5, 6, 7, 8, 9, 10).

func main() {
	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var value string = ""
	for _, number := range numbers {
		value += number
		fmt.Printf("%s\n", value)

	}

}
