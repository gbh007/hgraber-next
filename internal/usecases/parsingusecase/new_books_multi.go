package parsingusecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
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
		for _, agent := range agents {
			checkRes, err := uc.agentSystem.BooksCheck(ctx, agent.ID, []url.URL{multiURL})
			if len(checkRes) != 1 || !checkRes[0].IsPossible {
				continue
			}

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
					ctx, "agent check url",
					slog.String("agent_id", agent.ID.String()),
					slog.String("error", err.Error()),
					slog.String("url", multiURL.String()),
				)

				continue
			}

			list, err := uc.agentSystem.ParseList(ctx, agent.ID, multiURL, false)
			if err != nil {
				uc.logger.ErrorContext(
					ctx, "agent parse list",
					slog.String("agent_id", agent.ID.String()),
					slog.String("error", err.Error()),
					slog.String("url", multiURL.String()),
				)

				continue
			}

			if len(list.Books) == 0 {
				continue
			}

			newUrls := make([]url.URL, 0, len(list.Books))

			for _, info := range list.Books {
				u := info.ExtURL

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

				// Разбор деталей нужен только для проверки фильтра,
				// без него не тратим время на дорогой вызов.
				if len(flags.Filter.Fields.Attributes) > 0 || flags.Filter.Fields.Name != "" {
					pageCountLimit := 0

					extendInfo, err := uc.agentSystem.BookParse(ctx, agent.ID, u, &pageCountLimit)
					if err != nil {
						return parsing.MultiHandleMultipleResult{}, fmt.Errorf(
							"agent (%s) get book detail info (%s): %w", agent.ID.String(), u.String(), err,
						)
					}

					if !bookMatchesFilter(flags.Filter, extendInfo) {
						result.Details.RegisterSkipped(u)

						continue
					}
				}

				newUrls = append(newUrls, u)
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
