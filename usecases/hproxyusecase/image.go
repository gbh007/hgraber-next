package hproxyusecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) Image(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get agents: %w", err)
	}

	for _, agent := range agents {
		pagesInfo, err := uc.agentSystem.PagesCheck(
			ctx, agent.ID,
			[]agentmodel.AgentPageURL{
				{
					BookURL:  bookURL,
					ImageURL: imageURL,
				},
			},
		)

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
				ctx, "agent check page",
				slog.String("agent_id", agent.ID.String()),
				slog.String("error", err.Error()),
			)

			continue
		}

		for _, info := range pagesInfo {
			if !info.IsPossible || info.ImageURL.String() != imageURL.String() {
				continue
			}

			r, err := uc.agentSystem.PageLoad(ctx, agent.ID, agentmodel.AgentPageURL{
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
