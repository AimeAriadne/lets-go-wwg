package main

import (
	"fmt"
)

// crie uma slice de tamanho 12 usando a função make() e atribua os meses do ano a cada um dos elemntos individualmente.
// printe na tela essa slice mostrando todos os seus elementos.
// printe também uma fatia dessa slice do índice 2 ao 8.

func main() {
	meses := make([]string, 12)
	meses[0] = "Janeiro"
	meses[1] = "Fevereiro"
	meses[2] = "Março"
	meses[3] = "Abril"
	meses[4] = "Maio"
	meses[5] = "Junho"
	meses[6] = "Julho"
	meses[7] = "Agosto"
	meses[8] = "Setembro"
	meses[9] = "Outubro"
	meses[10] = "Novembro"
	meses[11] = "Dezembro"
	fmt.Println(meses)

	fmt.Println(meses[2:8])

}
