package webserver

import (
	"cep-retriever/internal/domain/entity/vo"
	"cep-retriever/internal/infra/service"
	"cep-retriever/internal/infra/service/impl"
	usecase "cep-retriever/internal/usecase/retrieve/impl"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
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
	apiCepCh := make(chan usecase.OutputDto)
	viaCepCh := make(chan usecase.OutputDto)

	cepCode := chi.URLParam(r, "cep")

	apicep := impl.NewApiCep()
	apiCepUrl := "https://cdn.apicep.com/file/apicep/cepValue.json"
	go func() {
		err := retrieveData(apicep, apiCepUrl, cepCode, apiCepCh)
		if err != nil {
			if err == vo.ErrCodeIsNotValid {
				w.WriteHeader(http.StatusBadRequest)
				errJson := Error{Message: err.Error()}
				err := json.NewEncoder(w).Encode(errJson)
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
		}
	}()

	viacep := impl.NewViaCep()
	viaCepurl := "https://viacep.com.br/ws/cepValue/json"
	go func() {
		err := retrieveData(viacep, viaCepurl, cepCode, apiCepCh)
		if err != nil {
			if err == vo.ErrCodeIsNotValid {
				w.WriteHeader(http.StatusBadRequest)
				errJson := Error{Message: err.Error()}
				err := json.NewEncoder(w).Encode(errJson)
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
		}
	}()

	select {
	case msg := <-apiCepCh:
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	case msg := <-viaCepCh:
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	case <-time.After(time.Second * 2):
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("no provider's response at minimum time")
		return
	}

}

func retrieveData(service service.Service, url string, cepCode string, outputCh chan usecase.OutputDto) error {
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
