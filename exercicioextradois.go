package main

import (
	"fmt"
	"time"
)

//1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
//2) Como podemos pegar a informação do ano diretamente do console?

func main() {
	var birthDate int
	currentYear := time.Now().Year()
	fmt.Println("Qual seu ano de nascimento?")
	fmt.Scan(&birthDate)
	age := currentYear - birthDate
	fmt.Printf("Sua idade é igual a: %v", age)
}
