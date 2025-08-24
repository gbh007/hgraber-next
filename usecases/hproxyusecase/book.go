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
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (uc *UseCase) Book(ctx context.Context, u url.URL, pageLimit *int) (hproxymodel.Book, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse:  true,
		HasHProxy: true,
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

			book, err := uc.agentSystem.HProxyBook(ctx, agent.ID, u, pageLimit)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("parse book: %w", err)
			}

			newAttrs, err := uc.handleAttributes(ctx, book.Attributes)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("handle attributes: %w", err)
			}

			book.Attributes = newAttrs

			if book.PreviewURL == nil {
				for _, page := range book.Pages {
					if page.PageNumber == core.PageNumberForPreview {
						book.PreviewURL = &page.ExtPreviewURL

						break
					}
				}
			}

			if book.PageCount == 0 {
				book.PageCount = len(book.Pages)
			}

			mirrors, err := uc.storage.Mirrors(ctx)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("get mirrors: %w", err)
			}

			mirrorCalculator := parsing.NewUrlCloner(mirrors)

			duplicates, err := mirrorCalculator.GetClones(u)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("calc duplicates: %w", err)
			}

			duplicates = append(duplicates, u)

			ids, err := uc.existsInStorage(ctx, duplicates)
			if err != nil {
				return hproxymodel.Book{}, fmt.Errorf("check duplicates: %w", err)
			}

			book.ExistsIDs = ids

			return book, nil
		}
	}

	return hproxymodel.Book{}, errCanNotParse
}
