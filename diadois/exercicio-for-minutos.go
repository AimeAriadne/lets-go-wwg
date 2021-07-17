package main

import (
	"fmt"
)

// utilizando a estrutura for, demostre as horas e os minutos de um dia (formato 24 horas)

func main() {
	hora := 0
	for hora < 24 {
		for minuto := 0; minuto < 60; minuto++ {
			fmt.Printf("%.2d:%.2d\n", hora, minuto)
		}
		hora++
	}

}
