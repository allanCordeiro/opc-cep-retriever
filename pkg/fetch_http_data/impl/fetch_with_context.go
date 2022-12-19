package impl

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type Fetch struct {
	body []byte
	err  error
}

func New() *Fetch {
	return &Fetch{}
}

func (f *Fetch) WithContext(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + res.Status)
	}

	f.body, f.err = io.ReadAll(res.Body)
	if f.err != nil {
		return nil, f.err
	}
	return f.body, nil
}
