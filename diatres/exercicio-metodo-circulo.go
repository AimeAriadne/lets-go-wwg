package main

import (
	"fmt"
	"math"
)

// Construa dois métodos: um que retorna a área e outro que retorna o perímetro de uma estrutura que representa um círculo. Printe a área e o perímetro de um círculo.
type Circulo struct {
	raio float64
}

func (c Circulo) CalculeArea() float64 {
	return math.Pi * (c.raio * c.raio)
}

func (c Circulo) CalculePerimetro() float64 {
	return 2 * math.Pi * c.raio
}
func main() {
	circulo1 := Circulo{
		raio: 6.5,
	}
	area := circulo1.CalculeArea()

	perimetro := circulo1.CalculePerimetro()

	fmt.Printf("círculo:\n\tárea: %.4f\n\tperímetro: %.4f\n\traio: %.4f", area, perimetro, circulo1.raio)

}
