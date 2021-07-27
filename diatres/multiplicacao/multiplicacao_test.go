package multiplicacao

import "testing"

func TestMultiplique(t *testing.T) {
	// chamar a função que multiplica
	obtive := Multiplique(10, 5)
	// dizer o que esperamos que aconteça
	espero := 50
	// comparar o que recebemos com o que esperávamos receber
	if obtive != espero {
		// sinalizar que o teste falhou
		t.Errorf("espero '%d', mas obtive '%d'", espero, obtive)
	}
}
