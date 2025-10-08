package parsingusecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

//nolint:gocognit,cyclop // будет исправлено позднее
func (uc *UseCase) NewBooksMulti(
	ctx context.Context,
	urls []url.URL,
	flags parsing.ParseFlags,
) (parsing.MultiHandleMultipleResult, error) {
	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParseMulti: true,
	})
	if err != nil {
		return parsing.MultiHandleMultipleResult{}, fmt.Errorf("get agents for parse: %w", err)
	}

	mirrors, err := uc.storage.Mirrors(ctx)
	if err != nil {
		return parsing.MultiHandleMultipleResult{}, fmt.Errorf("get mirrors: %w", err)
	}

	mirrorCalculator := parsing.NewURLCloner(mirrors)

	result := parsing.MultiHandleMultipleResult{
		Details: parsing.FirstHandleMultipleResult{
			Details: make([]parsing.BookHandleResult, 0, len(urls)*100), //nolint:mnd // оптимизация
		},
	}

urlLoop:
	for _, multiURL := range urls {
	agentLoop:
		for _, agent := range agents {
			booksInfo, err := uc.agentSystem.BooksCheckMultiple(ctx, agent.ID, multiURL)

			if errors.Is(err, agentmodel.ErrAgentAPIOffline) {
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

					urlsToCheck, err := mirrorCalculator.GetClones(u)
					if err != nil {
						return parsing.MultiHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) calc duplicates (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					urlsToCheck = append(urlsToCheck, u)

					ids, err := uc.existsInStorage(ctx, urlsToCheck)
					if err != nil {
						return parsing.MultiHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) check duplicates (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					if len(ids) > 0 {
						result.Details.RegisterDuplicate(u, ids)

						continue
					}

					newUrls = append(newUrls, u)
				}
			}

			for _, u := range newUrls {
				book := core.Book{
					ID:        uuid.Must(uuid.NewV7()),
					OriginURL: &u,
					CreateAt:  time.Now(),
				}

				if flags.AutoVerify {
					book.Verified = true
					book.VerifiedAt = time.Now().UTC()
				}

				if !flags.ReadOnly {
					err = uc.storage.NewBook(ctx, book)
					if err != nil {
						return parsing.MultiHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) create (%s): %w", agent.ID.String(), u.String(), err,
						)
					}
				}

				result.Details.RegisterHandled(u, book.ID)
			}

			result.RegisterHandled(multiURL)

			continue urlLoop
		}

		result.RegisterError(multiURL, "unsupported by all agents")
	}

	return result, nil
}
