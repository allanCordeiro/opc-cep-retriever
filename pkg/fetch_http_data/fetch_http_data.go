package fetch_http_data

import "context"

type FetchData interface {
	WithContext(ctx context.Context, url string) ([]byte, error)
}

func FetchWithContext(ctx context.Context, url string, fetch FetchData) ([]byte, error) {
	return fetch.WithContext(ctx, url)
}
