package agenthandlers

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
)

type AgentUseCases interface {
	NewAgent(ctx context.Context, agent core.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, filter core.AgentFilter, includeStatus bool) ([]core.AgentWithStatus, error)
	UpdateAgent(ctx context.Context, agent core.Agent) error
	Agent(ctx context.Context, id uuid.UUID) (core.Agent, error)
}

type ExportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, filter core.BookFilter, deleteAfter bool) error
}

type AgentHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	agentUseCases  AgentUseCases
	exportUseCases ExportUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	agentUseCases AgentUseCases,
	exportUseCases ExportUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *AgentHandlersController {
	c := &AgentHandlersController{
		logger:         logger,
		tracer:         tracer,
		agentUseCases:  agentUseCases,
		exportUseCases: exportUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}
