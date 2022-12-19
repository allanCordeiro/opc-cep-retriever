package impl

import (
	"cep-retriever/internal/infra/service"
	"cep-retriever/pkg/fetch_http_data"
	fetch "cep-retriever/pkg/fetch_http_data/impl"
	"context"
	"encoding/json"
	"strings"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep() *ViaCep {
	return &ViaCep{}
}

func (vc *ViaCep) RetrieveData(ctx context.Context, input service.Input) (service.Output, error) {
	url := input.Url
	url = strings.Replace(url, "cepValue", input.CepCode, 1)

	fetchImpl := fetch.New()
	body, err := fetch_http_data.FetchWithContext(ctx, url, fetchImpl)
	if err != nil {
		return service.Output{}, err
	}
	err = json.Unmarshal(body, &vc)
	if err != nil {
		return service.Output{}, err
	}

	return service.Output{
		Code:     vc.Cep,
		Address:  vc.Logradouro,
		District: vc.Bairro,
		City:     vc.Localidade,
		State:    vc.Uf,
	}, nil
}
