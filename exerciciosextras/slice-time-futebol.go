package main

import (
	"fmt"
)

// Existem dois times de futebol, o time amarelo e o time vermelho. O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana) e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
// crie uma slice para representar cada um.
// Adicione o jogador Luis Inácio no time vermelho usando a função append()
// Printe na tela os nomes dos jogadores do time vermelho

func main() {
	var timeAmarelo = []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}

	var timeVermelho = []string{"Helena", "Jonas", "José", "Juliana"}
	timeVermelho = append(timeVermelho, "Luis Inácio")

	fmt.Printf("Os jogadores do time amarelo são %s\n", timeAmarelo)
	fmt.Printf("Os jogadores do time vermelho são %s", timeVermelho)

}
