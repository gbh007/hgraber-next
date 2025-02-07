package parsing

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) NewBooks(ctx context.Context, urls []url.URL, autoVerify bool) (parsing.FirstHandleMultipleResult, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return parsing.FirstHandleMultipleResult{}, fmt.Errorf("get agents for parse: %w", err)
	}

	result := parsing.FirstHandleMultipleResult{
		Details: make([]parsing.BookHandleResult, 0, len(urls)),
	}

	urlSet := pkg.SliceToSet(urls)
	bookIDsByURL := pkg.SliceToMap(urls, func(u url.URL) (url.URL, uuid.UUID) {
		return u, uuid.Must(uuid.NewV7())
	})

	// Предварительная обработка, для уменьшения трафика на агенты
	for _, u := range urls {
		// Ссылки не могут содержать пробелы
		if u.String() != strings.TrimSpace(u.String()) {
			return parsing.FirstHandleMultipleResult{}, fmt.Errorf("url (%s) have space", u.String())
		}

		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return parsing.FirstHandleMultipleResult{}, fmt.Errorf("url exists in storage check (%s): %w", u.String(), err)
		}

		if len(ids) == 0 {
			continue
		}

		result.RegisterDuplicate(u)
		delete(urlSet, u)
	}

	for _, agent := range agents {
		if len(urlSet) == 0 {
			break
		}

		booksInfo, err := uc.agentSystem.BooksCheck(ctx, agent.ID, pkg.SetToSlice(urlSet))

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

		for _, info := range booksInfo {
			u := info.URL

			switch {
			case info.IsUnsupported:
				continue

			case info.HasError:
				result.RegisterError(u, info.ErrorReason)

			case info.IsPossible:
				if len(info.PossibleDuplicates) > 0 {
					exists, err := uc.existsInStorage(ctx, info.PossibleDuplicates)
					if err != nil {
						return parsing.FirstHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) check duplicates (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					if exists {
						result.RegisterDuplicate(u)

						break
					}
				}

				id, ok := bookIDsByURL[u]
				if !ok {
					uc.logger.WarnContext(
						ctx, "missing pregenerated book id",
						slog.String("book_url", u.String()),
					)

					continue
				}

				book := core.Book{
					ID:        id,
					OriginURL: &u,
					CreateAt:  time.Now().UTC(),
				}

				if autoVerify {
					book.Verified = true
					book.VerifiedAt = time.Now().UTC()
				}

				err = uc.storage.NewBook(ctx, book)
				if err != nil {
					return parsing.FirstHandleMultipleResult{}, fmt.Errorf(
						"agent (%s) create (%s): %w", agent.ID.String(), u.String(), err,
					)
				}

				result.RegisterHandled(u)
			}

			delete(urlSet, u)
		}
	}

	for u := range urlSet {
		result.RegisterError(u, "unsupported by all agents")
	}

	if len(result.Details) != len(urls) || result.TotalCount != len(urls) {
		uc.logger.WarnContext(
			ctx, "handled count not equivalent urls count",
			slog.Int("details_count", len(result.Details)),
			slog.Int("urls_count", len(urls)),
			slog.Int("total_count", result.TotalCount),
		)
	}

	return result, nil
}

func (uc *UseCase) existsInStorage(ctx context.Context, urls []url.URL) (bool, error) {
	for _, u := range urls {
		// FIXME: нужно сделать более оптимальный метод
		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return false, fmt.Errorf("check exists by (%s): %w", u.String(), err)
		}

		if len(ids) > 0 {
			return true, nil
		}
	}

	return false, nil
}
