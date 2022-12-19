package impl

import (
	"cep-retriever/internal/domain/entity"
	"cep-retriever/internal/domain/entity/vo"
	"cep-retriever/internal/infra/service"
	"context"
)

type UseCase struct {
	Cep     InputDto
	Url     string
	Ctx     context.Context
	service service.Service
}

type InputDto struct {
	Code string
}

type OutputDto struct {
	Code     string `json:"cep"`
	Address  string `json:"address"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
}

func NewUseCase(cep string, url string, ctx context.Context, svc service.Service) *UseCase {
	return &UseCase{
		Cep:     InputDto{cep},
		Url:     url,
		Ctx:     ctx,
		service: svc,
	}
}

func (uc *UseCase) Execute() (OutputDto, error) {
	code, err := vo.NewCode(uc.Cep.Code)
	if err != nil {
		return OutputDto{}, err
	}
	input := service.Input{CepCode: code.ID, Url: uc.Url}

	output, err := uc.service.RetrieveData(uc.Ctx, input)
	if err != nil {
		return OutputDto{}, err
	}

	domain, err := entity.NewCep(
		output.Code,
		output.Address,
		output.District,
		output.City,
		output.State,
	)

	if err != nil {
		return OutputDto{}, err
	}

	return OutputDto{
		Code:     domain.Code.ID,
		Address:  domain.Address,
		District: domain.District,
		City:     domain.City,
		State:    domain.State,
	}, nil
}
