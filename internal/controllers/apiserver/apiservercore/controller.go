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
	GetExternalAddr() string
	GetDebug() bool
	GetUseHeaderExternalAddr() bool
}

type Controller struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	fsUseCases FSUseCases

	externalServerScheme       string
	externalServerHostWithPort string
	useHeaderExternalAddr      bool
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	fsUseCases FSUseCases,
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
		debug:                      config.GetDebug(),
		useHeaderExternalAddr:      config.GetUseHeaderExternalAddr(),
	}

	return c, nil
}
