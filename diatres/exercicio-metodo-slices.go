package main

import (
	"fmt"
)

// Considere um cenário em que você precise regularmente trabalhar com slices de inteiros e extrair a soma e média dos números dessa lista. Usando o que você aprendeu sobre métodos, qual seria sua solução para esse problema em Go?

type numeros []int

func main() {
	var conjunto = numeros{1, 2, 3, 4, 5, 6}

	soma := conjunto.SomeNumeros()
	media := conjunto.CalculeMedia()

	fmt.Printf("a soma dos números presentes no conjunto é: %d\na média dos números presentes no conjunto é: %v", soma, media)

}
func (n numeros) SomeNumeros() int {
	var soma int = 0
	for _, numero := range n {
		soma += numero
	}
	return soma
}

func (n numeros) CalculeMedia() float64 {
	soma := float64(n.SomeNumeros())
	quantidade := float64(len(n))
	media := soma / quantidade
	return media
}
