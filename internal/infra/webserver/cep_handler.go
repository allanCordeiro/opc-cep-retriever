package webserver

import (
	"cep-retriever/internal/domain/entity/vo"
	"cep-retriever/internal/infra/service"
	"cep-retriever/internal/infra/service/impl"
	usecase "cep-retriever/internal/usecase/retrieve/impl"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type Provider struct {
	providerSvc service.Service
	providerUrl string
	ch          chan usecase.OutputDto
}

// CepHandleGet Find CEP through different providers
// @Summary     HandleGet
// @Description Find CEP through different providers
// @Tags        cep retriever
// @Produce     json
// @Param       cep path string true "cep code"
// @Success     200
// @Failure		400 {object} Error
// @Failure     500 {object} Error
// @Router      /retrieve/{cep} [get]
func CepHandleGet(w http.ResponseWriter, r *http.Request) {
	cepCode := chi.URLParam(r, "cep")
	g := new(errgroup.Group)
	providers := []Provider{
		{
			providerSvc: impl.NewApiCep(),
			providerUrl: "https://cdn.apicep.com/file/apicep/cepValue.json",
			ch:          make(chan usecase.OutputDto),
		},
		{
			providerSvc: impl.NewViaCep(),
			providerUrl: "https://viacep.com.br/ws/cepValue/json",
			ch:          make(chan usecase.OutputDto),
		},
	}

	for _, p := range providers {
		g.Go(func() error {
			err := getCepInfo(p.providerSvc, p.providerUrl, cepCode, p.ch)
			if err != nil {
				if err == vo.ErrCodeIsNotValid {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					errJson := Error{Message: err.Error()}
					err := json.NewEncoder(w).Encode(errJson)
					if err != nil {
						log.Println(err)
						return err
					}
					return err
				}
			}
			return nil
		})

		select {
		case msg := <-p.ch:
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(msg)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			return
		case <-time.After(time.Second * 2):
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("no provider's response at minimum time")
			return
		}
	}

}

func getCepInfo(service service.Service, url string, cepCode string, outputCh chan usecase.OutputDto) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	uc := usecase.NewUseCase(cepCode, url, ctx, service)
	output, err := uc.Execute()
	if err != nil {
		return err
	}
	outputCh <- output
	return nil
}
