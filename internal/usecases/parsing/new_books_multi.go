package parsing

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) NewBooksMulti(ctx context.Context, urls []url.URL, autoVerify bool) (entities.MultiHandleMultipleResult, error) {
	agents, err := uc.storage.Agents(ctx, entities.AgentFilter{
		CanParseMulti: true,
	})
	if err != nil {
		return entities.MultiHandleMultipleResult{}, fmt.Errorf("get agents for parse: %w", err)
	}

	result := entities.MultiHandleMultipleResult{
		Details: entities.FirstHandleMultipleResult{
			Details: make([]entities.BookHandleResult, 0, len(urls)*100),
		},
	}

urlLoop:
	for _, multiUrl := range urls {
	agentLoop:
		for _, agent := range agents {
			booksInfo, err := uc.agentSystem.BooksCheckMultiple(ctx, agent.ID, multiUrl)

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

			if len(booksInfo) == 0 {
				continue agentLoop
			}

			newUrls := make([]url.URL, 0, len(booksInfo))

			for _, info := range booksInfo {
				switch {
				case info.IsUnsupported:
					continue agentLoop

				case info.HasError:
					continue agentLoop

				case info.IsPossible:
					u := info.URL
					urlsToCheck := []url.URL{u}
					urlsToCheck = append(urlsToCheck, info.PossibleDuplicates...)

					exists, err := uc.existsInStorage(ctx, urlsToCheck)
					if err != nil {
						return entities.MultiHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) check duplicates (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					if exists {
						result.Details.RegisterDuplicate(u)

						continue
					}

					newUrls = append(newUrls, u)
				}
			}

			for _, u := range newUrls {
				book := entities.Book{
					ID:        uuid.Must(uuid.NewV7()),
					OriginURL: &u,
					CreateAt:  time.Now(),
				}

				if autoVerify {
					book.Verified = true
					book.VerifiedAt = time.Now().UTC()
				}

				err = uc.storage.NewBook(ctx, book)
				if err != nil {
					return entities.MultiHandleMultipleResult{}, fmt.Errorf(
						"agent (%s) create (%s): %w", agent.ID.String(), u.String(), err,
					)
				}

				result.Details.RegisterHandled(u)
			}

			result.RegisterHandled(multiUrl)

			continue urlLoop
		}

		result.RegisterError(multiUrl, "unsupported by all agents")
	}

	return result, nil
}
