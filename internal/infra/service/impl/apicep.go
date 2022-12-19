package impl

import (
	"cep-retriever/internal/infra/service"
	"cep-retriever/pkg/fetch_http_data"
	fetch "cep-retriever/pkg/fetch_http_data/impl"
	"context"
	"encoding/json"
	"strings"
)

type ApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func NewApiCep() *ApiCep {
	return &ApiCep{}
}

func (ac *ApiCep) RetrieveData(ctx context.Context, input service.Input) (service.Output, error) {
	url := input.Url
	url = strings.Replace(url, "cepValue", input.CepCode, 1)

	fetchImpl := fetch.New()
	body, err := fetch_http_data.FetchWithContext(ctx, url, fetchImpl)
	if err != nil {
		return service.Output{}, err
	}
	err = json.Unmarshal(body, &ac)
	if err != nil {
		return service.Output{}, err
	}

	return service.Output{
		Code:     ac.Code,
		Address:  ac.Address,
		District: ac.District,
		City:     ac.City,
		State:    ac.State,
	}, nil
}
