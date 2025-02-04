package agentcache

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
)

func (uc *UseCase) ParseBook(ctx context.Context, u url.URL) (agentmodel.AgentBookDetails, error) {
	book, err := uc.parseUseCases.BookByURL(ctx, u)
	if err != nil {
		return agentmodel.AgentBookDetails{}, err
	}

	return agentmodel.BookContainerToAgentBookDetails(book), nil
}
