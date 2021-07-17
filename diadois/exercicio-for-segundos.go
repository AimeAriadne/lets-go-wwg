package main

import (
	"fmt"
)

// utilizando a estrutura for, demostre as horas e os minutos de um dia e os segundos até 02:59:59.

func main() {
	hora := 0
	for hora <= 2 {
		for minuto := 0; minuto < 60; minuto++ {
			for segundo := 0; segundo < 60; segundo++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundo)
			}
		}
		hora++
	}

}
