package common

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/hgraber"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

// Проверка соответствия базового типа
var (
	_ hgraber.BookParser = (*BookParser)(nil)
)

type BookParser struct{}

func (BookParser) BookDetails(ctx context.Context) (entities.AgentBookDetails, error) {
	return entities.AgentBookDetails{}, nil
}

func (BookParser) HProxyBookDetails(ctx context.Context, pageLimit *int) (entities.HProxyBookDetails, error) {
	return entities.HProxyBookDetails{}, nil
}
