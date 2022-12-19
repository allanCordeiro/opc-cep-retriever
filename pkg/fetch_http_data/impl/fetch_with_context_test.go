package impl

import (
	"bytes"
	"cep-retriever/pkg/fetch_http_data"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type FetchMock struct {
	mock.Mock
}

func (m *FetchMock) WithContext(ctx context.Context, url string) ([]byte, error) {
	args := m.Called(ctx, url)
	return args.Get(0).([]byte), args.Error(1)
}

func TestGivenProperlyDataWhenCallFetchMethodThenShouldReturnData(t *testing.T) {
	expectedUrl := "https://www.someurl.com"
	expectedReturn := []byte("testing")
	expectedCtx := context.Background()
	expectedCtx, cancel := context.WithTimeout(expectedCtx, time.Millisecond*300)
	defer cancel()

	mocked := &FetchMock{}
	mocked.On("WithContext", expectedCtx, expectedUrl).Return(expectedReturn, nil)

	body, err := fetch_http_data.FetchWithContext(expectedCtx, expectedUrl, mocked)
	if err != nil {
		t.Errorf("an unexpected error was ocurred")
	}
	if !bytes.Equal(body, expectedReturn) {
		t.Errorf("expected '%s', but received '%s'", expectedReturn, body)
	}
}

func TestGivenProperlyDataWhenCallFetchAndAnErrorCorrursThenShouldReturnAnError(t *testing.T) {
	expectedUrl := "https://www.someurl.com"
	expectedCtx := context.Background()
	expectedCtx, cancel := context.WithTimeout(expectedCtx, time.Millisecond*300)
	defer cancel()

	mocked := &FetchMock{}
	mocked.On("WithContext", expectedCtx, expectedUrl).Return([]byte("0"), errors.New("http error"))

	_, err := fetch_http_data.FetchWithContext(expectedCtx, expectedUrl, mocked)
	if err == nil {
		t.Errorf("an error was expected but not happened")
	}

}
