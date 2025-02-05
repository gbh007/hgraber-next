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
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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
}

type Controller struct {
	*agenthandlers.AgentHandlersController
	*attributehandlers.AttributeHandlersController
	*bookhandlers.BookHandlersController
	*deduplicatehandlers.DeduplicateHandlersController
	*fshandlers.FSHandlersController
	*labelhandlers.LabelHandlersController
	*systemhandlers.SystemHandlersController

	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	ogenServer *serverAPI.Server

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
	debug bool,
) (*Controller, error) {
	ac, err := apiservercore.New(
		logger,
		tracer,
		config,
		fsUseCases,
		debug,
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
			debug,
			ac,
		),
		AttributeHandlersController: attributehandlers.New(
			logger,
			tracer,
			webAPIUseCases,
			debug,
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
			debug,
			ac,
		),
		DeduplicateHandlersController: deduplicatehandlers.New(
			logger,
			tracer,
			webAPIUseCases,
			deduplicateUseCases,
			debug,
			ac,
		),
		FSHandlersController: fshandlers.New(
			logger,
			tracer,
			parseUseCases,
			webAPIUseCases,
			taskUseCases,
			fsUseCases,
			debug,
			ac,
		),
		LabelHandlersController: labelhandlers.New(
			logger,
			tracer,
			webAPIUseCases,
			debug,
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
			debug,
			ac,
		),

		logger:     logger,
		tracer:     tracer,
		serverAddr: config.GetAddr(),
		debug:      debug,
		staticDir:  config.GetStaticDir(),
		token:      config.GetToken(),
	}

	ogenServer, err := serverAPI.NewServer(
		c, c,
		serverAPI.WithErrorHandler(methodErrorHandler),
		serverAPI.WithMethodNotAllowed(methodNotAllowed),
		serverAPI.WithNotFound(methodNotFound),
		serverAPI.WithMiddleware(c.simplePanicRecover),
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}
