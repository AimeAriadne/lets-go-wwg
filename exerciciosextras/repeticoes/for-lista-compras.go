package main

import (
	"fmt"
)

// Exercício 6: Dada uma slice de string que representa uma lista de mercado, use o for range para percorrê-la e printe na tela cada um dos itens de compra ao lado da sua ordem de inserção na lista.
// No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava uma lista de itens a comprar no mercado. Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for.

func main() {
	var listadeMercado = []string{"abacate", "sabonete", "azeite", "tomate", "banana"}

	for i := 0; i < 5; i++ {
		fmt.Println(listadeMercado[i])
	}

}
