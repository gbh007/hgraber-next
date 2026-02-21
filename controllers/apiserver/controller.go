//revive:disable:file-length-limit
package apiserver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/agenthandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/controllers/apiserver/attributehandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/bookhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/deduplicatehandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/fshandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/hproxyhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/labelhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/massloadhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/systemhandlers"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

type ParseUseCases interface {
	bookhandlers.ParseUseCases
	fshandlers.ParseUseCases
	systemhandlers.ParseUseCases
}

type BookUseCases interface {
	bookhandlers.BookUseCases
}

type LabelUseCases interface {
	labelhandlers.LabelUseCases
}

type AgentUseCases interface {
	agenthandlers.AgentUseCases
}

type ExportUseCases interface {
	agenthandlers.ExportUseCases
	bookhandlers.ExportUseCases
	systemhandlers.ExportUseCases
}

type DeduplicateUseCases interface {
	deduplicatehandlers.DeduplicateUseCases
	bookhandlers.DeduplicateUseCases
}

type SystemUseCases interface {
	fshandlers.SystemUseCases
	systemhandlers.SystemUseCases
}

type ReBuilderUseCases interface {
	bookhandlers.ReBuilderUseCases
}

type FSUseCases interface {
	apiservercore.FSUseCases
	fshandlers.FSUseCases
}

type BFFUseCases interface {
	bookhandlers.BFFUseCases
	deduplicatehandlers.BFFUseCases
}

type AttributeUseCases interface {
	attributehandlers.AttributeUseCases
}

type HProxyUseCases interface {
	hproxyhandlers.HProxyUseCases
}

type MassloadUseCases interface {
	massloadhandlers.MassloadUseCases
}

type config interface {
	GetAddr() string
	GetExternalAddr() string
	GetStaticDir() string
	GetToken() string
	GetLogErrorHandler() bool
	GetDebug() bool
}

type metricProvider interface {
	HTTPServerAddHandle(addr, operation string, status bool, d time.Duration)
	HTTPServerIncActive(addr, operation string)
	HTTPServerDecActive(addr, operation string)
	Registry() *prometheus.Registry
}

type Controller struct {
	*agenthandlers.AgentHandlersController
	*attributehandlers.AttributeHandlersController
	*bookhandlers.BookHandlersController
	*deduplicatehandlers.DeduplicateHandlersController
	*fshandlers.FSHandlersController
	*labelhandlers.LabelHandlersController
	*systemhandlers.SystemHandlersController
	*hproxyhandlers.HProxyHandlersController
	*massloadhandlers.MassloadController

	logger          *slog.Logger
	tracer          trace.Tracer
	metricProvider  metricProvider
	debug           bool
	logErrorHandler bool

	ogenServer *serverapi.Server

	staticDir  string
	serverAddr string
	token      string
}

func New( //revive:disable:argument-limit // будет исправлено позднее
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	metricProvider metricProvider,
	parseUseCases ParseUseCases,
	agentUseCases AgentUseCases,
	exportUseCases ExportUseCases,
	deduplicateUseCases DeduplicateUseCases,
	systemUseCases SystemUseCases,
	reBuilderUseCases ReBuilderUseCases,
	fsUseCases FSUseCases,
	bffUseCases BFFUseCases,
	attributeUseCases AttributeUseCases,
	labelUseCases LabelUseCases,
	bookUseCases BookUseCases,
	hProxyUseCases HProxyUseCases,
	massloadUseCases MassloadUseCases,
) (*Controller, error) {
	ac, err := apiservercore.New(
		logger,
		tracer,
		config,
		fsUseCases,
		config.GetDebug(),
	)
	if err != nil {
		return nil, fmt.Errorf("init core handlers: %w", err)
	}

	c := &Controller{
		AgentHandlersController: agenthandlers.New(
			logger,
			tracer,
			agentUseCases,
			exportUseCases,
			config.GetDebug(),
			ac,
		),
		AttributeHandlersController: attributehandlers.New(
			logger,
			tracer,
			attributeUseCases,
			config.GetDebug(),
			ac,
		),
		BookHandlersController: bookhandlers.New(
			logger,
			tracer,
			parseUseCases,
			bookUseCases,
			exportUseCases,
			reBuilderUseCases,
			bffUseCases,
			deduplicateUseCases,
			config.GetDebug(),
			ac,
		),
		DeduplicateHandlersController: deduplicatehandlers.New(
			logger,
			tracer,
			bffUseCases,
			deduplicateUseCases,
			config.GetDebug(),
			ac,
		),
		FSHandlersController: fshandlers.New(
			logger,
			tracer,
			parseUseCases,
			systemUseCases,
			fsUseCases,
			config.GetDebug(),
			ac,
		),
		LabelHandlersController: labelhandlers.New(
			logger,
			tracer,
			labelUseCases,
			config.GetDebug(),
			ac,
		),
		SystemHandlersController: systemhandlers.New(
			logger,
			tracer,
			parseUseCases,
			exportUseCases,
			systemUseCases,
			config.GetDebug(),
			ac,
		),
		HProxyHandlersController: hproxyhandlers.New(
			logger,
			tracer,
			hProxyUseCases,
			config.GetDebug(),
			ac,
		),
		MassloadController: massloadhandlers.New(
			logger,
			tracer,
			config.GetDebug(),
			ac,
			massloadUseCases,
		),

		logger:          logger,
		tracer:          tracer,
		metricProvider:  metricProvider,
		serverAddr:      config.GetAddr(),
		debug:           config.GetDebug(),
		logErrorHandler: config.GetLogErrorHandler(),
		staticDir:       config.GetStaticDir(),
		token:           config.GetToken(),
	}

	ogenServer, err := serverapi.NewServer(
		c, c,
		serverapi.WithErrorHandler(c.methodErrorHandler),
		serverapi.WithMethodNotAllowed(methodNotAllowed),
		serverapi.WithNotFound(methodNotFound),
		serverapi.WithMiddleware(
			c.metricsMiddleware,
			c.simplePanicRecover,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}

func (c *Controller) NewError(ctx context.Context, err error) *serverapi.ErrorResponseStatusCode {
	if err == nil {
		return &serverapi.ErrorResponseStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: serverapi.ErrorResponse{
				InnerCode: "unexpected",
				Details:   serverapi.NewOptString("missing error"),
			},
		}
	}

	var ae apiservercore.APIError

	if errors.As(err, &ae) {
		return &serverapi.ErrorResponseStatusCode{
			StatusCode: ae.Code,
			Response: serverapi.ErrorResponse{
				InnerCode: ae.InnerCode,
				Details: serverapi.OptString{
					Value: ae.Details,
					Set:   len(ae.Details) > 0,
				},
			},
		}
	}

	var (
		httpCode         = http.StatusInternalServerError
		errorCode        = "internal error"
		errorDescription = err.Error()
	)

	validateError := new(validate.Error)

	switch {
	case errors.Is(err, ogenerrors.ErrSecurityRequirementIsNotSatisfied):
		httpCode = http.StatusUnauthorized
		errorCode = "unauthorized"
	case errors.Is(err, errAccessForbidden):
		httpCode = http.StatusForbidden
		errorCode = "forbidden"
	case errors.Is(err, errPanicDetected):
		httpCode = http.StatusInternalServerError
		errorCode = "panic"
	case errors.As(err, &validateError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	}

	return &serverapi.ErrorResponseStatusCode{
		StatusCode: httpCode,
		Response: serverapi.ErrorResponse{
			InnerCode: errorCode,
			Details:   serverapi.NewOptString(errorDescription),
		},
	}
}
