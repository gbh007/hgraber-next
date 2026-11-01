package agentparseruc

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
)

func (uc *UseCase) Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error) {
	book, err := uc.loader.Book(ctx, u, pageLimit)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("load parser: %w", err)
	}

	return book, nil
}
