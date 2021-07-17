package main

import (
	"fmt"
)

// dada uma slice de string que representa uma lista de mercado, use o for range para percorrê-la e printe na tela cada um dos itens de compra ao lado da sua ordem de inserção na lista.

func main() {
	var listadeMercado = []string{"abacate", "sabonete", "azeite", "tomate", "banana"}

	for index, value := range listadeMercado {
		fmt.Printf("%d) - %s\n", index+1, value)
	}

}
