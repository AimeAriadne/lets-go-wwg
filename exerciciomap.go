package main

import (
	"fmt"
)

// crie um map chamado ano onde as chaves(key) são os números dos meses e o valor é o nome do mês.
// printe os meses 1 e 12.
// printe o tamanho do map ano.

func main() {
	mesesDoAno := make(map[int]string)
	mesesDoAno[1] = "Jan"
	mesesDoAno[2] = "Fev"
	mesesDoAno[3] = "Mar"
	mesesDoAno[4] = "Abr"
	mesesDoAno[5] = "Mai"
	mesesDoAno[6] = "Jun"
	mesesDoAno[7] = "Jul"
	mesesDoAno[8] = "Ago"
	mesesDoAno[9] = "Set"
	mesesDoAno[10] = "Out"
	mesesDoAno[11] = "Nov"
	mesesDoAno[12] = "Dez"

	fmt.Println(mesesDoAno[1], mesesDoAno[12])
	fmt.Println(len(mesesDoAno))
}
