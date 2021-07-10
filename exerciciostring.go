package main

import (
	"fmt"
)

// declare duas variáveis do tipo string, uma vai receber seu nome a outra vai receber sua cor favorita.
// printe na tela uma frase utilizando os dois valores.
func main() {
	name := "Aimê"
	favoriteColor := "black"
	fmt.Printf("My name is %s and my favorite color is %s.", name, favoriteColor)
}

// soliticando informação do usuário
// func main() {
//	var name string
//	var favoriteColor string
//	fmt.Println("What is your name?")
//	fmt.Scan(&name)
//	fmt.Println("What is your favorite color?")
//	fmt.Scan(&favoriteColor)
//	fmt.Printf("Your name is %s and your favorite color is %s", name, favoriteColor)
//}
