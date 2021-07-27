package main

import (
	"fmt"
	"math"
)

//Crie tipos para representar quadrados e círculos
//Crie uma interface que descreve o comportamento de calcular a área de uma forma geométrica com a seguinte assinatura: calculeArea() float64
//Implemente esse comportamento para os dois tipos criados
//Depois, crie uma função que tem como parâmetro a interface que você criou e que imprime o relatório do cálculo da área da forma geométrica
//Demonstre que seus tipos implementam a interface que você criou passando valores desses tipos como argumentos na chamada dessa função

type Formas interface {
	calculeArea() float64
}

type Quadrados struct {
	lado float64
}

func (q Quadrados) calculeArea() float64 {
	area := (q.lado * q.lado)
	return area
}

type Circulos struct {
	raio float64
}

func (c Circulos) calculeArea() float64 {
	area := math.Pi * (c.raio * c.raio)
	return area
}

func reportarCalculo(forma Formas) {
	area := forma.calculeArea()
	fmt.Printf("a área da forma em questão é %.4f metros\n", area)
}

func main() {
	quadrado1 := Quadrados{
		lado: 13,
	}
	circulo1 := Circulos{
		raio: 27,
	}

	reportarCalculo(quadrado1)
	reportarCalculo(circulo1)
}
