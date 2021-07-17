package main

import (
	"fmt"
)

// utilizando a estrutura for, printe na tela os todas as horas do dia(usando o formato de 24 horas)

func main() {
	hora := 0
	for hora < 24 {
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}

}
