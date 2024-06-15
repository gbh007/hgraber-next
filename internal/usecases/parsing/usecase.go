package parsing

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type Storage interface {
	NewBook(ctx context.Context, book entities.Book) error
	GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error)
	Agents(ctx context.Context, canParse, canExport bool) ([]entities.Agent, error)
}

type AgentSystem interface {
	BookParse(ctx context.Context, agentID uuid.UUID, url url.URL) (entities.AgentBookDetails, error)
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	PageLoad(ctx context.Context, agentID uuid.UUID, url entities.AgentPageURL) (io.Reader, error)
	PagesCheck(ctx context.Context, agentID uuid.UUID, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
}

type UseCase struct {
	logger *slog.Logger

	storage     Storage
	agentSystem AgentSystem
}

func New(
	logger *slog.Logger,
	storage Storage,
	agentSystem AgentSystem,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		agentSystem: agentSystem,
	}
}
