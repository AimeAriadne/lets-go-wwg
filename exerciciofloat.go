package main

import (
	"fmt"
)

// calcure o valor total de uma compra que tem 3 itens representando o preço de todos eles em floar64.
// todos os itens dessa compra precisam ser comprados em mais de uma unidade.

func main() {
	var chocolate float64 = 7.99
	var pipoca float64 = 3.29
	var doritos float64 = 6.99
	var compras float64 = (chocolate * 2) + (pipoca * 3) + (doritos * 2)
	fmt.Printf("O valor da compra deu %v, só uma pipoca já custou %v, só um chocolate custou %v, só um doritos custou %v\n", compras, pipoca, chocolate, doritos)
}

// outra forma
func main2() {
	var tomate, azeite, sabonete float64
	tomate = 5.55
	azeite = 10.99
	sabonete = 1.29
	var total float64 = tomate + azeite + sabonete
	fmt.Println(total)
}
