package main

import (
	"fmt"
)

// dado um número A e um número B, escreva uma função que some ambos e retorne a soma.

func main() {
	soma := Some(10, 5)
	fmt.Println(soma)
}

func Some(a, b int) int {
	return a + b
}
