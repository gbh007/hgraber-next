package parsing

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) BooksToParse(ctx context.Context) ([]entities.BookWithAgent, error) {
	books, err := uc.storage.UnprocessedBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("books from storage: %w", err)
	}

	if len(books) == 0 {
		return []entities.BookWithAgent{}, nil
	}

	agents, err := uc.storage.Agents(ctx, entities.AgentFilter{
		CanParse: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get agents: %w", err)
	}

	books = pkg.SliceFilter(books, func(b entities.Book) bool {
		hasUrl := b.OriginURL != nil
		if !hasUrl {
			uc.logger.WarnContext(
				ctx, "handle book without url",
				slog.String("book_id", b.ID.String()),
			)
		}

		return hasUrl
	})

	toParse := make([]entities.BookWithAgent, 0, len(books))

	urlMap := pkg.SliceToMap(books, func(b entities.Book) (url.URL, entities.Book) {
		return *b.OriginURL, b
	})

	for _, agent := range agents {
		if len(urlMap) == 0 {
			break
		}

		booksInfo, err := uc.agentSystem.BooksCheck(ctx, agent.ID, pkg.MapToSlice(urlMap, func(_ url.URL, b entities.Book) url.URL {
			return *b.OriginURL
		}))

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
			if !info.IsPossible {
				continue
			}

			book, ok := urlMap[info.URL]
			if !ok {
				continue
			}

			toParse = append(toParse, entities.BookWithAgent{
				Book:    book,
				AgentID: agent.ID,
			})

			delete(urlMap, info.URL)
		}
	}

	if len(toParse) != len(books) || len(urlMap) > 0 {
		uc.logger.WarnContext(
			ctx, "handled count not equivalent books count",
			slog.Int("to_parse_count", len(toParse)),
			slog.Int("book_count", len(books)),
			slog.Int("left_book_count", len(urlMap)),
		)
	}

	return toParse, nil
}
