package aniversario

import (
	"testing"

	"github.com/jeysonlr/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaAniversarioFantasma(t *testing.T) {
	s := NewAniversarioInfantil()
	
	_, err := s.Calcula(domain.Parametros{
		Homens:          0,
		Mulheres:        1,
		Criancas:        2,
		Acompanhamentos: true,
	})
	assert.Equal(t, err.Error(), "Homens, mulheres ou criaças devem ser maiores que zero")
}

func TestCalculaAniversarioBombando(t *testing.T) {
	s := NewAniversarioInfantil()
	
	ch, err := s.Calcula(domain.Parametros{
		Homens:          5,
		Mulheres:        5,
		Criancas:        10,
		Acompanhamentos: false,
	})

	assert.Nil(t, err)

	esperado := Churrasco{
		TotalPessoas:         20,
		Refrigerante:         5000,
		Salgadinhos:          150,
		Brigadeiros:          150,
		Bolo:                 2000,
	}
	assert.Equal(t, ch, esperado)
}
