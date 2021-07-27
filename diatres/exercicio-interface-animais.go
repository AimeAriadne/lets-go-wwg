package main

import (
	"fmt"
)

//Crie tipos que representam diferentes animais, com atributos que façam sentido para cada um deles
//Crie uma interface que descreve o comportamento de apresentar um animal com a seguinte assinatura: Apresenta()
//Cada animal saberá como se apresentar. Sendo assim, faça com que cada um dos tipos que você criou implemente o método Apresenta(), que deve printar uma frase apresentando o animal e seus atributos
//Demonstre que todos os tipos implementam a interface que você criou declarando uma slice de animais e percorrendo-a com um for range que, em todas as voltas, chama o método Apresenta().
type Animal interface {
	Apresenta()
}

type Cachorro struct {
	nome string
	raca string
}

func (c Cachorro) Apresenta() {
	fmt.Printf("Au au! Meu nome é %s e eu sou um %s.\n", c.nome, c.raca)
}

type Gato struct {
	nome      string
	corDoPelo string
}

func (g Gato) Apresenta() {
	fmt.Printf("Miau! Meu nome é %s e a cor do meu pelo é %s.\n", g.nome, g.corDoPelo)
}

type Peixe struct {
	nome string
	cor  string
}

func (p Peixe) Apresenta() {
	fmt.Printf("Glub glub! Meu nome é %s e eu sou %s.\n", p.nome, p.cor)
}

func main() {
	cachorro1 := Cachorro{
		nome: "Duque",
		raca: "Fila",
	}

	gato1 := Gato{
		nome:      "Satanás",
		corDoPelo: "preto",
	}

	peixe1 := Peixe{
		nome: "Peixonauta",
		cor:  "dourado",
	}

	animais := []Animal{cachorro1, gato1, peixe1}

	for _, animal := range animais {
		animal.Apresenta()
	}
}
