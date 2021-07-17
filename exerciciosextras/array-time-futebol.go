package main

import (
	"fmt"
)

// Existem dois times de futebol, o time amarelo e o time vermelho. O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana) e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
// Crie um array de string para cada time e nomeie com o nome do time.
// Printe na tela os nomes dos jogadores de cada time

func main() {
	var timeAmarelo [5]string
	timeAmarelo[0] = "Fernando"
	timeAmarelo[1] = "João"
	timeAmarelo[2] = "Lúcia"
	timeAmarelo[3] = "Mariana"
	timeAmarelo[4] = "Ana"

	var timeVermelho [4]string
	timeVermelho[0] = "Helena"
	timeVermelho[1] = "Jonas"
	timeVermelho[2] = "José"
	timeVermelho[3] = "Juliana"

	fmt.Printf("Os jogadores do time amarelo são %s\n", timeAmarelo)
	fmt.Printf("Os jogadores do time vermelho são %s", timeVermelho)

}
