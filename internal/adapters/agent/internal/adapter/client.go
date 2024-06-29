package adapter

import (
	"net/http"
	"time"

	"hgnext/internal/adapters/agent/internal/client"
)

const agentTimeout = time.Minute * 10

type Adapter struct {
	rawClient *client.Client
}

// TODO: возможно стоит вынести инициализацию HTTP клиента наружу
func New(baseURL string, token string) (*Adapter, error) {
	httpClient := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   agentTimeout,
	}

	rawClient, err := client.NewClient(
		baseURL,
		securitySource{
			token: token,
		},
		client.WithClient(&httpClient),
	)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		rawClient: rawClient,
	}, nil
}
