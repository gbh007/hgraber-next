package parsingusecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"net/url"
	"slices"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/samber/lo"
)

func (uc *UseCase) PagesToDownload(ctx context.Context) ([]parsing.PageForDownloadWithAgent, error) {
	pages, err := uc.storage.NotDownloadedPages(ctx)
	if err != nil {
		return nil, fmt.Errorf("pages from storage: %w", err)
	}

	if len(pages) == 0 {
		return []parsing.PageForDownloadWithAgent{}, nil
	}

	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get agents: %w", err)
	}

	pages = lo.Filter(pages, func(b core.PageForDownload, _ int) bool {
		hasURL := b.BookURL != nil && b.ImageURL != nil
		if !hasURL {
			uc.logger.WarnContext(
				ctx, "handle page without url",
				slog.String("book_id", b.BookID.String()),
				slog.Int("page_number", b.PageNumber),
			)
		}

		return hasURL
	})

	toDownload := make([]parsing.PageForDownloadWithAgent, 0, len(pages))

	urlMap := make(map[url.URL][]core.PageForDownload)

	for _, p := range pages {
		urlMap[*p.BookURL] = append(urlMap[*p.BookURL], p)
	}

	for _, agent := range agents {
		if len(urlMap) == 0 {
			break
		}

		pagesInfo, err := uc.agentSystem.BooksCheck(
			ctx, agent.ID,
			slices.Collect(maps.Keys(urlMap)),
		)

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

			urlPages, ok := urlMap[info.URL]
			if !ok {
				continue
			}

			for _, page := range urlPages {
				toDownload = append(toDownload, parsing.PageForDownloadWithAgent{
					PageForDownload: page,
					AgentID:         agent.ID,
				})
			}

			delete(urlMap, info.URL)
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
