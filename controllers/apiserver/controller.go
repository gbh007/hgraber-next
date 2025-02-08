package apiserver

import (
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/agenthandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/controllers/apiserver/attributehandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/bookhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/deduplicatehandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/fshandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/labelhandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/systemhandlers"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

type ParseUseCases interface {
	bookhandlers.ParseUseCases
	fshandlers.ParseUseCases
	systemhandlers.ParseUseCases
}

type WebAPIUseCases interface {
	attributehandlers.WebAPIUseCases
	bookhandlers.WebAPIUseCases
	deduplicatehandlers.WebAPIUseCases
	fshandlers.WebAPIUseCases
	labelhandlers.WebAPIUseCases
	systemhandlers.WebAPIUseCases
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
	systemhandlers.DeduplicateUseCases
}

type TaskUseCases interface {
	fshandlers.TaskUseCases
	systemhandlers.TaskUseCases
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
}

type config interface {
	GetAddr() string
	GetExternalAddr() string
	GetStaticDir() string
	GetToken() string
	GetLogErrorHandler() bool
	GetDebug() bool
}

type Controller struct {
	*agenthandlers.AgentHandlersController
	*attributehandlers.AttributeHandlersController
	*bookhandlers.BookHandlersController
	*deduplicatehandlers.DeduplicateHandlersController
	*fshandlers.FSHandlersController
	*labelhandlers.LabelHandlersController
	*systemhandlers.SystemHandlersController

	logger          *slog.Logger
	tracer          trace.Tracer
	debug           bool
	logErrorHandler bool

	ogenServer *serverapi.Server

	staticDir  string
	serverAddr string
	token      string
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	parseUseCases ParseUseCases,
	webAPIUseCases WebAPIUseCases,
	agentUseCases AgentUseCases,
	exportUseCases ExportUseCases,
	deduplicateUseCases DeduplicateUseCases,
	taskUseCases TaskUseCases,
	rebuilderUseCases ReBuilderUseCases,
	fsUseCases FSUseCases,
	bffUseCases BFFUseCases,
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
			webAPIUseCases,
			config.GetDebug(),
			ac,
		),
		BookHandlersController: bookhandlers.New(
			logger,
			tracer,
			parseUseCases,
			webAPIUseCases,
			exportUseCases,
			rebuilderUseCases,
			bffUseCases,
			config.GetDebug(),
			ac,
		),
		DeduplicateHandlersController: deduplicatehandlers.New(
			logger,
			tracer,
			webAPIUseCases,
			deduplicateUseCases,
			config.GetDebug(),
			ac,
		),
		FSHandlersController: fshandlers.New(
			logger,
			tracer,
			parseUseCases,
			webAPIUseCases,
			taskUseCases,
			fsUseCases,
			config.GetDebug(),
			ac,
		),
		LabelHandlersController: labelhandlers.New(
			logger,
			tracer,
			webAPIUseCases,
			config.GetDebug(),
			ac,
		),
		SystemHandlersController: systemhandlers.New(
			logger,
			tracer,
			parseUseCases,
			webAPIUseCases,
			exportUseCases,
			deduplicateUseCases,
			taskUseCases,
			config.GetDebug(),
			ac,
		),

		logger:          logger,
		tracer:          tracer,
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
		serverapi.WithMiddleware(c.simplePanicRecover),
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}
