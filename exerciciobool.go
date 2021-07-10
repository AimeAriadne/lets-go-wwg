package main

import (
	"fmt"
)

// declare 5 variáveis e atribua operações relacionais a elas.
// usando uma linha por variável, demonstre o valor e o tipo de cada uma delas.
func main() {
	a := 120 >= 120
	b := 150 < 120 || 120 > 100
	c := 5 != 5
	d := 50 != 100
	e := 90 > 10 && 50 > 15

	fmt.Printf("O tipo de a é: %T e seu valor é %v\n", a, a)
	fmt.Printf("O tipo de b é: %T e seu valor é %v\n", b, b)
	fmt.Printf("O tipo de c é: %T e seu valor é %v\n", c, c)
	fmt.Printf("O tipo de d é: %T e seu valor é %v\n", d, d)
	fmt.Printf("O tipo de e é: %T e seu valor é %v\n", e, e)

}
