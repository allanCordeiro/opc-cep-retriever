package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type RetrieveDataMock struct {
	mock.Mock
}

func (m *RetrieveDataMock) RetrieveData(ctx context.Context, input Input) (Output, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(Output), args.Error(1)
}

func TestNewApiCep(t *testing.T) {
	input := Input{CepCode: "01234-567", Url: "https:www.someurl.com"}
	expectedOutput := Output{
		Code:     "01211-100",
		Address:  "Some address",
		District: "Some district",
		City:     "Some city",
		State:    "SA",
	}
	expectedCtx := context.Background()
	expectedCtx, cancel := context.WithTimeout(expectedCtx, time.Millisecond*300)
	defer cancel()
	mocked := &RetrieveDataMock{}
	mocked.On("RetrieveData", expectedCtx, input).Return(expectedOutput, nil)

	output, err := mocked.RetrieveData(expectedCtx, input)
	if err != nil {
		t.Errorf("an unexpected error was ocurred")
	}

	if output.Code != expectedOutput.Code {
		t.Errorf("expected '%s', but received '%s'", expectedOutput.Code, output.Code)
	}
	if output.Address != expectedOutput.Address {
		t.Errorf("expected '%s', but received '%s'", expectedOutput.Address, output.Address)
	}
	if output.District != expectedOutput.District {
		t.Errorf("expected '%s', but received '%s'", expectedOutput.District, output.District)
	}
	if output.City != expectedOutput.City {
		t.Errorf("expected '%s', but received '%s'", expectedOutput.City, output.City)
	}
	if output.State != expectedOutput.State {
		t.Errorf("expected '%s', but received '%s'", expectedOutput.State, output.State)
	}
}
