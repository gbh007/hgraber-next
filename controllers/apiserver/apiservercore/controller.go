package apiservercore

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type FSUseCases interface {
	HighwayFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) (url.URL, bool, error)
}

type config interface {
	GetAddr() string
	GetExternalAddr() string
	GetStaticDir() string
	GetToken() string
}

type Controller struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	fsUseCases FSUseCases

	externalServerScheme       string
	externalServerHostWithPort string
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	fsUseCases FSUseCases,
	debug bool,
) (*Controller, error) {
	u, err := url.Parse(config.GetExternalAddr())
	if err != nil {
		return nil, fmt.Errorf("parse external server addr: %w", err)
	}

	c := &Controller{
		logger:                     logger,
		tracer:                     tracer,
		externalServerScheme:       u.Scheme,
		externalServerHostWithPort: u.Host,
		fsUseCases:                 fsUseCases,
		debug:                      debug,
	}

	return c, nil
}
