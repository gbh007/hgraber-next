package masterapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/ogen-go/ogen/ogenerrors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const agentTimeout = time.Minute * 1

type Client struct {
	rawClient *serverapi.Client
}

func New(baseURL string, token string) (*Client, error) {
	httpClient := http.Client{
		Transport: otelPropagationRT{next: http.DefaultTransport},
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

type otelPropagationRT struct {
	next http.RoundTripper
}

func (rt otelPropagationRT) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))

	return rt.next.RoundTrip(req)
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
