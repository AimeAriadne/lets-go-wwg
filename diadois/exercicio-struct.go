package main

import (
	"fmt"
)

//Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
//Crie três variáveis do tipo pessoa.
//Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.

type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {
	bay := Pessoa{
		nome:         "Bay",
		idade:        18,
		corPreferida: "roxo"}
	cecil := Pessoa{
		nome:         "Cecil",
		idade:        20,
		corPreferida: "preto"}
	joy := Pessoa{
		nome:         "Joy",
		idade:        25,
		corPreferida: "verde"}

	fmt.Printf("%s tem %v anos e sua cor preferida é %s\n", bay.nome, bay.idade, bay.corPreferida)
	fmt.Printf("%s tem %v anos e sua cor preferida é %s\n", cecil.nome, cecil.idade, cecil.corPreferida)
	fmt.Printf("%s tem %v anos e sua cor preferida é %s\n", joy.nome, joy.idade, joy.corPreferida)

}
