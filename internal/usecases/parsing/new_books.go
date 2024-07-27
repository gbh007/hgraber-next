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

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) NewBooks(ctx context.Context, urls []url.URL) (entities.FirstHandleMultipleResult, error) {
	agents, err := uc.storage.Agents(ctx, true, false)
	if err != nil {
		return entities.FirstHandleMultipleResult{}, fmt.Errorf("get agents for parse: %w", err)
	}

	result := entities.FirstHandleMultipleResult{
		Details: make([]entities.BookHandleResult, 0, len(urls)),
	}

	urlSet := pkg.SliceToSet(urls)

	// Предварительная обработка, для уменьшения трафика на агенты
	for _, u := range urls {
		// Ссылки не могут содержать пробелы
		if u.String() != strings.TrimSpace(u.String()) {
			return entities.FirstHandleMultipleResult{}, fmt.Errorf("url (%s) have space", u.String())
		}

		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return entities.FirstHandleMultipleResult{}, fmt.Errorf("url exists in storage check (%s): %w", u.String(), err)
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
						return entities.FirstHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) check duplicates (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					if exists {
						result.RegisterDuplicate(u)

						break
					}
				}

				// TODO: потеря порядка, из-за множеств, возможно нет причины устранять
				err = uc.storage.NewBook(ctx, entities.Book{
					ID:        uuid.Must(uuid.NewV7()),
					OriginURL: &u,
					CreateAt:  time.Now(),
				})
				if err != nil {
					return entities.FirstHandleMultipleResult{}, fmt.Errorf(
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
