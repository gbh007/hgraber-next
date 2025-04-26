package hproxyusecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
)

func (uc *UseCase) Book(ctx context.Context, u url.URL) (hproxymodel.Book, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse:  true,
		CanHProxy: true,
	})
	if err != nil {
		return hproxymodel.Book{}, fmt.Errorf("get agents: %w", err)
	}

	for _, agent := range agents {
		info, err := uc.agentSystem.BooksCheck(ctx, agent.ID, []url.URL{u})

		if errors.Is(err, agentmodel.AgentAPIOffline) {
			uc.logger.DebugContext(
				ctx, "agent api offline",
				slog.String("agent_id", agent.ID.String()),
				slog.String("error", err.Error()),
			)

			continue
		}

		if err != nil {
			uc.logger.ErrorContext(
				ctx, "agent check book",
				slog.String("agent_id", agent.ID.String()),
				slog.String("error", err.Error()),
			)

			continue
		}

		for _, res := range info {
			if !res.IsPossible || res.URL.String() != u.String() {
				continue
			}

			list, err := uc.agentSystem.HProxyBook(ctx, agent.ID, u)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("parse book: %w", err)
			}

			return list, nil
		}
	}

	return hproxymodel.Book{}, errCanNotParse
}
