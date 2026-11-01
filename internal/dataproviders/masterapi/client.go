package masterapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/ogen-go/ogen/ogenerrors"
)

const agentTimeout = time.Minute * 1

type Client struct {
	rawClient *serverapi.Client
}

func New(baseURL, token string) (*Client, error) {
	httpClient := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   agentTimeout,
	}

	rawClient, err := serverapi.NewClient(
		baseURL,
		securitySource{
			token: token,
		},
		serverapi.WithClient(&httpClient),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		rawClient: rawClient,
	}, nil
}

type securitySource struct {
	token string
}

func (s securitySource) HeaderAuth(ctx context.Context, operationName string) (serverapi.HeaderAuth, error) {
	return serverapi.HeaderAuth{
		APIKey: s.token,
	}, nil
}

func (s securitySource) Cookies(ctx context.Context, operationName string) (serverapi.Cookies, error) {
	return serverapi.Cookies{}, ogenerrors.ErrSkipClientSecurity
}
