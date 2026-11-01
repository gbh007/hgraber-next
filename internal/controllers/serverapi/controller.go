package serverapi

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/agenthandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/attributehandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/bookhandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/deduplicatehandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/fshandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/hproxyhandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/labelhandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/massloadhandlers"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/systemhandlers"
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
	GetUseHeaderExternalAddr() bool
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
	metricProvider  metricProvider
	debug           bool
	logErrorHandler bool

	ogenServer *serverapi.Server

	staticDir             string
	serverAddr            string
	token                 string
	useHeaderExternalAddr bool
}

func New( //revive:disable:argument-limit // будет исправлено позднее
	logger *slog.Logger,
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
		config,
		fsUseCases,
	)
	if err != nil {
		return nil, fmt.Errorf("init core handlers: %w", err)
	}

	c := &Controller{
		AgentHandlersController: agenthandlers.New(
			logger,
			agentUseCases,
			exportUseCases,
			config.GetDebug(),
			ac,
		),
		AttributeHandlersController: attributehandlers.New(
			logger,
			attributeUseCases,
			config.GetDebug(),
			ac,
		),
		BookHandlersController: bookhandlers.New(
			logger,
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
			bffUseCases,
			deduplicateUseCases,
			config.GetDebug(),
			ac,
		),
		FSHandlersController: fshandlers.New(
			logger,
			parseUseCases,
			systemUseCases,
			fsUseCases,
			config.GetDebug(),
			ac,
		),
		LabelHandlersController: labelhandlers.New(
			logger,
			labelUseCases,
			config.GetDebug(),
			ac,
		),
		SystemHandlersController: systemhandlers.New(
			logger,
			parseUseCases,
			exportUseCases,
			systemUseCases,
			config.GetDebug(),
			ac,
		),
		HProxyHandlersController: hproxyhandlers.New(
			logger,
			hProxyUseCases,
			config.GetDebug(),
			ac,
		),
		MassloadController: massloadhandlers.New(
			logger,
			config.GetDebug(),
			ac,
			massloadUseCases,
		),

		logger:                logger,
		metricProvider:        metricProvider,
		serverAddr:            config.GetAddr(),
		debug:                 config.GetDebug(),
		logErrorHandler:       config.GetLogErrorHandler(),
		staticDir:             config.GetStaticDir(),
		token:                 config.GetToken(),
		useHeaderExternalAddr: config.GetUseHeaderExternalAddr(),
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
