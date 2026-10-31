package hproxyusecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

func (uc *UseCase) Image(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get agents: %w", err)
	}

	for _, agent := range agents {
		pagesInfo, err := uc.agentSystem.BooksCheck(ctx, agent.ID, []url.URL{bookURL})

		if errors.Is(err, serveragentdto.ErrAgentAPIOffline) {
			uc.logger.DebugContext(
				ctx, "agent api offline",
				slog.String("agent_id", agent.ID.String()),
				slog.String("error", err.Error()),
			)

			continue
		}

		if err != nil {
			uc.logger.ErrorContext(
				ctx, "agent check page",
				slog.String("agent_id", agent.ID.String()),
				slog.String("error", err.Error()),
			)

			continue
		}

		for _, info := range pagesInfo {
			if !info.IsPossible || info.URL.String() != bookURL.String() {
				continue
			}

			r, err := uc.agentSystem.PageLoad(ctx, agent.ID, serveragentdto.AgentPageURL{
				BookURL:  bookURL,
				ImageURL: imageURL,
			})
			if err != nil {
				return nil, fmt.Errorf("load page: %w", err)
			}

			return r, nil
		}
	}

	return nil, errCanNotParse
}
