package main

import "testing"

// SOMA correto
func TestSomaCorrect(t *testing.T) {
	teste := soma(2, 4, 6)
	resultado := 12
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// SOMA Incorreto
func TestSomaIncorrect(t *testing.T) {
	teste := soma(2, 4, 6)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// SUBTRACAO correto
func TestSubtracaoCorrect(t *testing.T) {
	teste := subtracao(10, 20)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// SUBTRACAO Incorreto
func TestSubtracaoIncorrect(t *testing.T) {
	teste := subtracao(10, 20)
	resultado := 30
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// MULTIPLICACAO Correto
func TestMultiplicacaoCorreto(t *testing.T) {
	teste := multiplicacao(100, 100)
	resultado := 10000
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// MULTIPLICACAO Incorreto
func TestMultiplicacaoIncorreto(t *testing.T) {
	teste := multiplicacao(100, 100)
	resultado := 1000
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// DIVISAO Correto
func TestDivisaoCorreto(t *testing.T) {
	teste := divisao(100)
	resultado := 50
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// DIVISAO Incorreto
func TestDivisaoIncorreto(t *testing.T) {
	teste := divisao(100)
	resultado := 20
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
