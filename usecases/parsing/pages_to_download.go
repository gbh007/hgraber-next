package parsing

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) PagesToDownload(ctx context.Context) ([]entities.PageForDownloadWithAgent, error) {
	pages, err := uc.storage.NotDownloadedPages(ctx)
	if err != nil {
		return nil, fmt.Errorf("pages from storage: %w", err)
	}

	if len(pages) == 0 {
		return []entities.PageForDownloadWithAgent{}, nil
	}

	agents, err := uc.storage.Agents(ctx, entities.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get agents: %w", err)
	}

	pages = pkg.SliceFilter(pages, func(b entities.PageForDownload) bool {
		hasUrl := b.BookURL != nil && b.ImageURL != nil
		if !hasUrl {
			uc.logger.WarnContext(
				ctx, "handle page without url",
				slog.String("book_id", b.BookID.String()),
				slog.Int("page_number", b.PageNumber),
			)
		}

		return hasUrl
	})

	toDownload := make([]entities.PageForDownloadWithAgent, 0, len(pages))

	urlMap := pkg.SliceToMap(pages, func(p entities.PageForDownload) (entities.AgentPageURL, entities.PageForDownload) {
		return entities.AgentPageURL{
			BookURL:  *p.BookURL,
			ImageURL: *p.ImageURL,
		}, p
	})

	for _, agent := range agents {
		if len(urlMap) == 0 {
			break
		}

		pagesInfo, err := uc.agentSystem.PagesCheck(
			ctx, agent.ID,
			pkg.MapToSlice(urlMap, func(u entities.AgentPageURL, _ entities.PageForDownload) entities.AgentPageURL {
				return u
			}),
		)

		if errors.Is(err, entities.AgentAPIOffline) {
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

		for _, info := range pagesInfo {
			if !info.IsPossible {
				continue
			}

			u := entities.AgentPageURL{
				BookURL:  info.BookURL,
				ImageURL: info.ImageURL,
			}

			page, ok := urlMap[u]
			if !ok {
				continue
			}

			toDownload = append(toDownload, entities.PageForDownloadWithAgent{
				PageForDownload: page,
				AgentID:         agent.ID,
			})

			delete(urlMap, u)
		}
	}

	if len(toDownload) != len(pages) || len(urlMap) > 0 {
		uc.logger.WarnContext(
			ctx, "handled count not equivalent pages count",
			slog.Int("to_download_count", len(toDownload)),
			slog.Int("page_count", len(pages)),
			slog.Int("left_page_count", len(urlMap)),
		)
	}

	return toDownload, nil
}
