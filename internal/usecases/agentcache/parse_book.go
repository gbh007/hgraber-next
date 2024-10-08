package agentcache

import (
	"context"
	"net/url"

	"hgnext/internal/entities"
)

func (uc *UseCase) ParseBook(ctx context.Context, u url.URL) (entities.AgentBookDetails, error) {
	book, err := uc.parseUseCases.BookByURL(ctx, u)
	if err != nil {
		return entities.AgentBookDetails{}, err
	}

	return book.ToAgentBookDetails(), nil
}
