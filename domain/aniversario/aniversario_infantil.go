package aniversario

import (
	"encoding/json"
	"errors"

	"github.com/jeysonlr/curso-go/domain"
)

const (
	qtdBoloPessoa = 100
	qdtSalgadinhoOuDoceAdulto = 10
	qtdSalgadinhoOuDoceCrianca = 5
	qtdRefrigerante = 500
)

type Churrasco struct {
	TotalPessoas         int `json:"total-pessoas"`
	Refrigerante         int `json:"refrigerante"`
	Salgadinhos          int `json:"salgadinhos"`
	Brigadeiros          int `json:"brigadeiros"`
	Bolo                 int `json:"bolo"`
}

func (c Churrasco) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

type Service struct{}

func NewAniversarioInfantil() Service {
	return Service{}
}

func (s Service) Calcula(request domain.Parametros) (domain.Resultado, error) {
	niver := Churrasco{}

	if request.Criancas == 0 || request.Homens == 0 || request.Mulheres == 0 {
		return niver, errors.New("Homens, mulheres ou criaças devem ser maiores que zero")
	}

	niver.TotalPessoas = request.Criancas + request.Mulheres + request.Homens
	niver.Bolo = request.Criancas * qtdBoloPessoa + request.Mulheres * qtdBoloPessoa + request.Homens * qtdBoloPessoa
	niver.Salgadinhos = request.Criancas * qtdSalgadinhoOuDoceCrianca + request.Mulheres * qdtSalgadinhoOuDoceAdulto + request.Homens * qdtSalgadinhoOuDoceAdulto
	niver.Brigadeiros = request.Criancas * qtdSalgadinhoOuDoceCrianca + request.Mulheres * qdtSalgadinhoOuDoceAdulto + request.Homens * qdtSalgadinhoOuDoceAdulto
	niver.Refrigerante = qtdRefrigerante * (request.Mulheres + request.Homens)

	return niver, nil
}

