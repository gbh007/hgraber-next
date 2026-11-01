package agenthandlers

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

type AgentUseCases interface {
	NewAgent(ctx context.Context, agent core.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, filter core.AgentFilter, includeStatus bool) ([]serveragentdto.AgentWithStatus, error)
	UpdateAgent(ctx context.Context, agent core.Agent) error
	Agent(ctx context.Context, id uuid.UUID) (core.Agent, error)
}

type ExportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, filter core.BookFilter, deleteAfter bool) error
}

type AgentHandlersController struct {
	logger *slog.Logger
	debug  bool

	apiCore *apiservercore.Controller

	agentUseCases  AgentUseCases
	exportUseCases ExportUseCases
}

func New(
	logger *slog.Logger,
	agentUseCases AgentUseCases,
	exportUseCases ExportUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *AgentHandlersController {
	c := &AgentHandlersController{
		logger:         logger,
		agentUseCases:  agentUseCases,
		exportUseCases: exportUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}
