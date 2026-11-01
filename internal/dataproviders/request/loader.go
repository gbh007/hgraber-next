package request

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Cache interface {
	SetString(ctx context.Context, u, v string)
	GetString(ctx context.Context, u string) (string, bool)
	SetBytes(ctx context.Context, u string, v []byte)
	GetBytes(ctx context.Context, u string) ([]byte, bool)
}

type Requester struct {
	logger           *slog.Logger
	cache            Cache
	client           *http.Client
	defaultUserAgent string
}

func New(
	logger *slog.Logger,
	timeout time.Duration,
	transport http.RoundTripper,
	logErrors bool,
	cache Cache,
	defaultUserAgent string,
) *Requester {
	if transport == nil {
		transport = http.DefaultTransport
	}

	if logErrors {
		transport = &requestErrorLogger{
			logger: logger,
			next:   transport,
		}
	}

	if defaultUserAgent == "" {
		defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
	}

	return &Requester{
		client: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		logger:           logger,
		cache:            cache,
		defaultUserAgent: defaultUserAgent,
	}
}

// RequestString запрашивает данные по урле и возвращает их строкой
func (r *Requester) RequestString(ctx context.Context, u string) (string, error) {
	if r.cache != nil {
		v, ok := r.cache.GetString(ctx, u)
		if ok {
			return v, nil
		}
	}

	buff, _, err := r.requestBuffer(ctx, u, nil, nil)
	if err != nil {
		return "", err
	}

	if r.cache != nil {
		r.cache.SetString(ctx, u, buff.String())
	}

	return buff.String(), nil
}

func (r *Requester) RequestStringWithRedirect(ctx context.Context, u string) (string, string, error) {
	if r.cache != nil {
		v, ok := r.cache.GetString(ctx, u)
		if ok {
			return v, u, nil
		}
	}

	buff, resultURL, err := r.requestBuffer(ctx, u, nil, nil)
	if err != nil {
		return "", "", err
	}

	if r.cache != nil {
		r.cache.SetString(ctx, resultURL, buff.String())
	}

	return buff.String(), resultURL, nil
}

// RequestBytes запрашивает данные по урле и возвращает их массивом байт
func (r *Requester) RequestBytes(ctx context.Context, u string) ([]byte, error) {
	buff, _, err := r.requestBuffer(ctx, u, nil, nil)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (r *Requester) Request(ctx context.Context, u string, headers http.Header) (io.ReadCloser, error) {
	if r.cache != nil {
		v, ok := r.cache.GetBytes(ctx, u)
		if ok {
			return io.NopCloser(bytes.NewReader(v)), nil
		}
	}

	// FIXME: работать с потоком напрямую
	buff, _, err := r.requestBuffer(ctx, u, headers, nil)
	if err != nil {
		return nil, err
	}

	if r.cache != nil {
		r.cache.SetBytes(ctx, u, buff.Bytes())
	}

	return io.NopCloser(buff), nil
}

func (r *Requester) RequestPost(ctx context.Context, u string, headers http.Header, body io.Reader) ([]byte, error) {
	buff, _, err := r.requestBuffer(ctx, u, headers, body)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
