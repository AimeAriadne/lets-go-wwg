package main

import "fmt"

// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
// Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
// Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse mapa segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
// Printe na tela os valores do último mapa.

func main() {
	listaDePaises := map[string]string{
		"Salvador":  "Brasil",
		"São Paulo": "Brasil",
		"Curitiba":  "Brasil",
		"São Luís":  "Brasil",
		"Londres":   "Inglaterra",
		"Cambridge": "Inglaterra",
		"Galway":    "Irlanda",
		"Dublin":    "Irlanda",
		"Cork":      "Irlanda",
	}

	frequencia := make(map[string]int)

	for _, value := range listaDePaises {
		frequencia[value] += 1
	}
	fmt.Println(frequencia)

}
