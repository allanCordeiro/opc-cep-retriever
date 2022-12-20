package impl

import (
	"cep-retriever/internal/infra/service"
	"context"
	"testing"
	"time"
)

type apiCepOutput struct {
	Code     string
	Address  string
	District string
	City     string
	State    string
}

func TestNewApiCep(t *testing.T) {
	expectedCepOutput := apiCepOutput{
		Code:     "01211-100",
		Address:  "Avenida São João - de 1341 a 1789 - lado ímpar",
		District: "Santa Cecília",
		City:     "São Paulo",
		State:    "SP",
	}

	t.Run("Given a valid data, when calls retrieve data, should return expected structure", func(t *testing.T) {
		input := service.Input{
			CepCode: "01211-100",
			Url:     "https://cdn.apicep.com/file/apicep/cepValue.json",
		}
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1000)
		defer cancel()

		apicep := NewApiCep()
		output, err := apicep.RetrieveData(ctx, input)
		if err != nil {
			t.Errorf("An unexpected error has been occurred")
		}

		if output.Code != expectedCepOutput.Code {
			t.Errorf("Data expected was '%s' but '%s' was found", expectedCepOutput.Code, output.Code)
		}

		if output.Address != expectedCepOutput.Address {
			t.Errorf("Data expected was '%s' but '%s' was found", expectedCepOutput.Address, output.Address)
		}

		if output.District != expectedCepOutput.District {
			t.Errorf("Data expected was '%s' but '%s' was found", expectedCepOutput.District, output.District)
		}

		if output.City != expectedCepOutput.City {
			t.Errorf("Data expected was '%s' but '%s' was found", expectedCepOutput.City, output.City)
		}

		if output.State != expectedCepOutput.State {
			t.Errorf("Data expected was '%s' but '%s' was found", expectedCepOutput.State, output.State)
		}
	})
}
