package adapter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

type Adapter struct {
	rawClient *agentapi.Client
}

type FSAdapter struct {
	rawClient *agentapi.Client
}

// TODO: возможно стоит вынести инициализацию HTTP клиента наружу
func New(baseURL, token string, agentTimeout time.Duration) (*Adapter, error) {
	httpClient := http.Client{
		Transport: agentOfflineRoundTripper{next: otelPropagationRoundTripper{next: http.DefaultTransport}},
		Timeout:   agentTimeout,
	}

	rawClient, err := agentapi.NewClient(
		baseURL,
		securitySource{
			token: token,
		},
		agentapi.WithClient(&httpClient),
	)
	if err != nil {
		return nil, fmt.Errorf("api new client: %w", err)
	}

	return &Adapter{
		rawClient: rawClient,
	}, nil
}

// TODO: возможно такая изоляция избыточна, и будет достаточно сделать 1 адаптер с уникальными названиями.
func (a *Adapter) ToFS() *FSAdapter {
	return &FSAdapter{
		rawClient: a.rawClient,
	}
}
