package massloadusecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

type handledExternalURL struct {
	urlsAhead    []url.URL
	urlsNew      []url.URL
	urlsExisting []url.URL
}

//nolint:cyclop // будет исправлено позднее
func (uc *UseCase) calcExternalLink(ctx context.Context, u url.URL) (handledExternalURL, error) {
	mirrors, err := uc.storage.Mirrors(ctx)
	if err != nil {
		return handledExternalURL{}, fmt.Errorf("get mirrors: %w", err)
	}

	mirrorCalculator := parsing.NewURLCloner(mirrors)

	agents, err := uc.storage.Agents(ctx, core.AgentFilter{
		CanParseMulti: true,
		HasHProxy:     true,
	})
	if err != nil {
		return handledExternalURL{}, fmt.Errorf("get agents: %w", err)
	}

	var agentID uuid.UUID

agentSearch:
	for _, agent := range agents {
		// TODO: чекаем по книге, хотя на самом деле здесь список,
		// это плохо, но в текущих реализациях парсеров будет работать
		info, err := uc.agentSystem.BooksCheck(ctx, agent.ID, []url.URL{u})

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

		for _, res := range info {
			if res.IsPossible && res.URL.String() == u.String() {
				agentID = agent.ID

				break agentSearch
			}
		}
	}

	if agentID == uuid.Nil {
		return handledExternalURL{}, errors.New("can't parse")
	}

	var (
		result   handledExternalURL
		foundAny bool
		nextURL  = &u
	)

	for nextURL != nil {
		if ctx.Err() != nil {
			return handledExternalURL{}, fmt.Errorf("ctx error: %w", ctx.Err())
		}

		uc.logger.DebugContext(
			ctx, "start calculate external url",
			slog.String("url", nextURL.String()),
		)

		list, err := uc.agentSystem.HProxyList(ctx, agentID, *nextURL)
		if err != nil {
			return handledExternalURL{}, fmt.Errorf("parse list: %w", err)
		}

		for _, book := range list.Books {
			duplicates, err := mirrorCalculator.GetClones(book.ExtURL)
			if err != nil {
				return handledExternalURL{}, fmt.Errorf("calc duplicates: %w", err)
			}

			duplicates = append(duplicates, book.ExtURL)

			ids, err := uc.existsInStorage(ctx, duplicates)
			if err != nil {
				return handledExternalURL{}, fmt.Errorf("check duplicates: %w", err)
			}

			switch {
			case len(ids) > 0:
				foundAny = true

				result.urlsExisting = append(result.urlsExisting, book.ExtURL)

			case foundAny:
				result.urlsNew = append(result.urlsNew, book.ExtURL)

			default:
				result.urlsNew = append(result.urlsNew, book.ExtURL)
				result.urlsAhead = append(result.urlsAhead, book.ExtURL)
			}
		}

		nextURL = list.NextPage
	}

	return result, nil
}

func (uc *UseCase) existsInStorage(ctx context.Context, urls []url.URL) ([]uuid.UUID, error) {
	for _, u := range urls {
		// TODO: нужно сделать более оптимальный метод
		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return nil, fmt.Errorf("check exists by (%s): %w", u.String(), err)
		}

		if len(ids) > 0 {
			return ids, nil
		}
	}

	return []uuid.UUID{}, nil
}
