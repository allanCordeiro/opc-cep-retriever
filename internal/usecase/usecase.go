package usecase

import "cep-retriever/internal/infra/service"

type UseCase interface {
	Execute(output service.Output, err error)
}
