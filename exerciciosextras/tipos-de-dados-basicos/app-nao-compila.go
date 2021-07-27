package main

import (
	"fmt"
)

//A aplicação abaixo não compila.
//package main
//import (
//	"fmt"
//)
//func main() {
//	var quilometros int8
//	quilometros = 150
//	fmt.Println(quilometros)
//}
//1) Descubra por que não compila
//2) Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
//3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela

func main() {
	var quilometros int
	quilometros = 150

	fmt.Println(quilometros)
}

// a aplicação anterior não compila, pois o tipo int8 aceita valores entre -128 ~ 127 e o valor atrbuído é 150, então overflows.
// resolve-se modificando o tipo para int ou modificando o valor atribuído para a faixa aceitada por int8.
