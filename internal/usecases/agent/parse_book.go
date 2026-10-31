package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

func (uc *UseCase) ParseBook(ctx context.Context, u url.URL) (entities.AgentBookDetails, error) {
	parser, err := uc.loader.Load(ctx, u)
	if err != nil {
		return entities.AgentBookDetails{}, fmt.Errorf("load parser: %w", err)
	}

	details, err := parser.BookDetails(ctx)
	if err != nil {
		return entities.AgentBookDetails{}, fmt.Errorf("parse: %w", err)
	}

	return details, nil
}
