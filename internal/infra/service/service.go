package service

import "context"

type Input struct {
	Url     string
	CepCode string
}

type Output struct {
	Code     string
	Address  string
	District string
	City     string
	State    string
}

type Service interface {
	RetrieveData(ctx context.Context, i Input) (Output, error)
}
