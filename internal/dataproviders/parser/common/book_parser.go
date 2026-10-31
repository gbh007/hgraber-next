package common

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

// Проверка соответствия базового типа
var (
	_ parsing.BookParser = (*BookParser)(nil)
)

type BookParser struct{}

func (BookParser) BookDetails(ctx context.Context) (serveragentdto.AgentBookDetails, error) {
	return serveragentdto.AgentBookDetails{}, nil
}

func (BookParser) HProxyBookDetails(ctx context.Context, pageLimit *int) (agentdto.HProxyBookDetails, error) {
	return agentdto.HProxyBookDetails{}, nil
}
