package adapter

import (
	"errors"
	"fmt"
	"net/http"
	"syscall"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

type Adapter struct {
	rawClient *agentAPI.Client
}

type FSAdapter struct {
	rawClient *agentAPI.Client
}

// TODO: возможно стоит вынести инициализацию HTTP клиента наружу
func New(baseURL string, token string, agentTimeout time.Duration) (*Adapter, error) {
	httpClient := http.Client{
		Transport: agentOfflineRT{next: otelPropagationRT{next: http.DefaultTransport}},
		Timeout:   agentTimeout,
	}

	rawClient, err := agentAPI.NewClient(
		baseURL,
		securitySource{
			token: token,
		},
		agentAPI.WithClient(&httpClient),
	)
	if err != nil {
		return nil, err
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

type otelPropagationRT struct {
	next http.RoundTripper
}

func (rt otelPropagationRT) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))

	return rt.next.RoundTrip(req)
}

type agentOfflineRT struct {
	next http.RoundTripper
}

func (rt agentOfflineRT) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := rt.next.RoundTrip(req)

	if errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EHOSTDOWN) ||
		errors.Is(err, syscall.EHOSTUNREACH) {
		err = fmt.Errorf("%w: %w", entities.AgentAPIOffline, err)
	}

	return res, err
}
